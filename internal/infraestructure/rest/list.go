package rest

import (
	"encoding/json"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"net/http"
)

type RESTListUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewRESTListUseCase(repository persistence.ReplacemeRepository) *RESTListUseCase {
	return &RESTListUseCase{
		repository: repository,
	}
}

func (uc *RESTListUseCase) List(w http.ResponseWriter, r *http.Request) {
	listUC := services.NewListReplacemeUseCase(uc.repository)
	replacemes, err := listUC.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := make([]replacemeDTO, 0, len(replacemes))
	for _, a := range replacemes {
		response = append(response, replacemeDTO{
			ID:        a.ID.Value(),
			Title:     a.Title.Value(),
			Content:   a.Content.Value(),
			Poster:    a.Poster.Value(),
			Tags:      a.Tags.Value(),
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
