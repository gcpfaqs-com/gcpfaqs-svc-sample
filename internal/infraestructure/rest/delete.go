package rest

import (
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"net/http"
)

type RESTDeleteUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewRESTDeleteUseCase(repository persistence.ReplacemeRepository) *RESTDeleteUseCase {
	return &RESTDeleteUseCase{
		repository: repository,
	}
}

func (uc *RESTDeleteUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	deleteUC := services.NewDeleteReplacemeUseCase(uc.repository)
	if err := deleteUC.Do(idStr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
