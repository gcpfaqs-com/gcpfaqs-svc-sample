package services

import (
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
)

type DeleteReplacemeUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewDeleteReplacemeUseCase(repo persistence.ReplacemeRepository) *DeleteReplacemeUseCase {
	return &DeleteReplacemeUseCase{repository: repo}
}

func (uc *DeleteReplacemeUseCase) Do(id string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	err = uc.repository.Delete(&validID)
	if err != nil {
		return err
	}

	return nil
}
