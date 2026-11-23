package rest

import (
	"encoding/json"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"net/http"
)

type RESTUpdateUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewRESTUpdateUseCase(repository persistence.ReplacemeRepository) *RESTUpdateUseCase {
	return &RESTUpdateUseCase{
		repository: repository,
	}
}

func (uc *RESTUpdateUseCase) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var dto replacemeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updateUC := services.NewUpdateReplacemeUseCase(uc.repository)
	err := updateUC.Do(
		idStr,
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

	dto.ID = idStr
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}
