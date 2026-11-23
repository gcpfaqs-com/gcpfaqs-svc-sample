package services

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"time"
)

type SaveReplacemeUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewSaveReplacemeUseCase(repo persistence.ReplacemeRepository) *SaveReplacemeUseCase {

	return &SaveReplacemeUseCase{repository: repo}

}

func (uc *SaveReplacemeUseCase) Do(
	title,
	content,
	poster string,
	tags []string,
	created_at, updated_at time.Time) (string, error) {
	replaceme, err := models.NewReplaceme(title, content, poster, tags, created_at, updated_at)
	if err != nil {
		return "", err
	}

	err = uc.repository.Save(replaceme)
	if err != nil {
		return "", err
	}

	return replaceme.ID.Value(), nil
}
