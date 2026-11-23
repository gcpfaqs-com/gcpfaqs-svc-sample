package services

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
)

type ListReplacemeUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewListReplacemeUseCase(repo persistence.ReplacemeRepository) *ListReplacemeUseCase {
	return &ListReplacemeUseCase{repository: repo}
}

func (uc *ListReplacemeUseCase) Do() ([]*models.Replaceme, error) {
	replacemes, err := uc.repository.List()
	if err != nil {
		return nil, err
	}

	return replacemes, nil
}
