package persistence

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
)

type ReplacemeRepository interface {
	Save(replaceme *models.Replaceme) error
	Get(id *types.ID) (*models.Replaceme, error)
	Delete(id *types.ID) error
	Update(replaceme *models.Replaceme) error
	List() ([]*models.Replaceme, error)
}
