package services

import (
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"time"
)

type UpdateReplacemeUseCase struct {
	repository persistence.ReplacemeRepository
}

func NewUpdateReplacemeUseCase(repo persistence.ReplacemeRepository) *UpdateReplacemeUseCase {

	return &UpdateReplacemeUseCase{repository: repo}

}

func (uc *UpdateReplacemeUseCase) Do(
	id,
	title,
	content,
	poster string,
	tags []string,
	createdAt, updatedAt time.Time) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	replaceme, err := uc.repository.Get(&validID)
	if err != nil {
		return err
	}

	validContent, err := types.NewContent(content)
	if err != nil {
		return err
	}

	validPoster, err := types.NewPoster(poster)
	if err != nil {
		return err
	}

	validTitle, err := types.NewTitle(title)
	if err != nil {
		return err
	}

	validTags, err := types.NewTags(tags)
	if err != nil {
		return err
	}

	replaceme.Title = validTitle
	replaceme.Content = validContent
	replaceme.Poster = validPoster
	replaceme.Tags = validTags
	replaceme.CreatedAt = createdAt
	replaceme.UpdatedAt = updatedAt
	if err := uc.repository.Update(replaceme); err == nil {
		return err
	}

	return nil
}
