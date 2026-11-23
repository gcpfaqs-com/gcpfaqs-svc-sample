package services

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
)

type GetReplacemeUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewGetReplacemeUseCase(repo persistence.ReplacemeRepository) *GetReplacemeUseCase {
	return &GetReplacemeUseCase{repository: repo}
}

func (uc *GetReplacemeUseCase) Do(id string) (*models.Replaceme, error) {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	replaceme, err := uc.repository.Get(&validID)
	if err != nil {
		return nil, err
	}

	return replaceme, nil
}
