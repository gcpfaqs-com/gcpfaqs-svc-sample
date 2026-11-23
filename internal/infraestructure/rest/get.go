package rest

import (
	"encoding/json"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"net/http"
)

type RESTGetUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewRESTGetUseCase(repository persistence.ReplacemeRepository) *RESTGetUseCase {
	return &RESTGetUseCase{
		repository: repository,
	}
}

func (uc *RESTGetUseCase) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	getUC := services.NewGetReplacemeUseCase(uc.repository)
	replaceme, err := getUC.Do(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := replacemeDTO{
		ID:        replaceme.ID.Value(),
		Title:     replaceme.Title.Value(),
		Content:   replaceme.Content.Value(),
		Poster:    replaceme.Poster.Value(),
		Tags:      replaceme.Tags.Value(),
		CreatedAt: replaceme.CreatedAt,
		UpdatedAt: replaceme.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
