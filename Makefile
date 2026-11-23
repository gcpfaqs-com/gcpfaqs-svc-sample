start-db:
	docker compose -f docker-compose.dev.yaml up -d

stop-db:
	docker compose -f docker-compose.dev.yaml down

start:
	SERVER_ADDRESS=":8081" DATABASE_DSN="host=localhost port=5431 user=postgres password=postgres dbname=gcpfaqs-svc-replacemes_test sslmode=disable" go run cmd/rest/main.go

test:
	go clean -testcache && go test  ./...