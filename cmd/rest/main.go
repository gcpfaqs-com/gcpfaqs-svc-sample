package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gcpreplacemes-svc-replacemes/internal/infraestructure/persistence/postgres"
	"gcpreplacemes-svc-replacemes/internal/infraestructure/rest"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const databaseDSNSecretID = "replacemes-database-dsn"

func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("database connectivity check failed: %v", err)
	}

	repo := postgres.NewPostgresReplacemeRepository(db)
	saveHandler := rest.NewRESTSaveUseCase(repo)
	getHandler := rest.NewRESTGetUseCase(repo)
	updateHandler := rest.NewRESTUpdateUseCase(repo)
	deleteHandler := rest.NewRESTDeleteUseCase(repo)
	listHandler := rest.NewRESTListUseCase(repo)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/replacemes", func(w http.ResponseWriter, r *http.Request) {
		listHandler.List(w, r)
	})

	mux.HandleFunc("GET /api/replacemes/{id}", func(w http.ResponseWriter, r *http.Request) {
		getHandler.Get(w, r)
	})

	mux.HandleFunc("POST /api/replacemes", func(w http.ResponseWriter, r *http.Request) {
		saveHandler.Save(w, r)

	})

	mux.HandleFunc("PUT /api/replacemes/{id}", func(w http.ResponseWriter, r *http.Request) {
		updateHandler.Update(w, r)
	})

	mux.HandleFunc("DELETE /api/replacemes/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteHandler.Delete(w, r)
	})

	mux.HandleFunc("GET /api/replacemes/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	address := resolveAddress()
	log.Printf("Database connectivity OK. Starting REST service on port %s...", address)
	if err := http.ListenAndServe(address, corsMiddleware(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func connectDatabase() (*gorm.DB, error) {
	dsn, err := resolveDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	return db, nil
}

func resolveDSN() (string, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = os.Getenv("GCP_PROJECT")
	}

	var secretValue string
	if projectID != "" {
		value, err := readDatabaseDSNSecret(projectID)
		if err != nil {
			return "", err
		}
		secretValue = value
	}

	return resolveDatabaseDSN(secretValue, nil, os.Getenv("DATABASE_DSN"))
}

func readDatabaseDSNSecret(projectID string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("create Secret Manager client: %w", err)
	}
	defer client.Close()

	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, databaseDSNSecretID),
	})
	if status.Code(err) == codes.NotFound {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("access Secret Manager secret %q: %w", databaseDSNSecretID, err)
	}
	if result == nil || result.GetPayload() == nil {
		return "", nil
	}

	return string(result.GetPayload().GetData()), nil
}

func resolveDatabaseDSN(secretValue string, secretErr error, environmentValue string) (string, error) {
	if secretErr != nil {
		return "", secretErr
	}
	if strings.TrimSpace(secretValue) != "" {
		return secretValue, nil
	}
	if environmentValue != "" {
		return environmentValue, nil
	}

	return "", errors.New("DATABASE_DSN must be set in Secret Manager secret replacemes-database-dsn or in the DATABASE_DSN environment variable")
}

func resolveAddress() string {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		log.Fatal("SERVER_ADDRESS must be set")
	}

	return address
}
