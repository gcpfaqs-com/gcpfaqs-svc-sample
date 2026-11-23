package rest

import (
	"encoding/json"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"net/http"
)

type RESTSaveUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewRESTSaveUseCase(repository persistence.ReplacemeRepository) *RESTSaveUseCase {
	return &RESTSaveUseCase{
		repository: repository,
	}
}

func (uc *RESTSaveUseCase) Save(w http.ResponseWriter, r *http.Request) {
	var dto replacemeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	saveUC := services.NewSaveReplacemeUseCase(uc.repository)
	id, err := saveUC.Do(
		dto.Title,
		dto.Content,
		dto.Poster,
		dto.Tags,
		dto.CreatedAt,
		dto.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto)
}
