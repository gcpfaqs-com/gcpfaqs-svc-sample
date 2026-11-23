package postgres

import (
	"errors"
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/persistence"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"time"

	"gorm.io/gorm"
)

type PostgresReplacemeRepository struct {
	db *gorm.DB
}

type replacemeRecord struct {
	ID        string    `gorm:"primaryKey;type:uuid;not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Poster    string    `gorm:"not null"`
	Tags      []string  `gorm:"serializer:json;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (replacemeRecord) TableName() string {
	return "replacemes"
}

func NewPostgresReplacemeRepository(db *gorm.DB) persistence.ReplacemeRepository {
	if db == nil {
		panic("postgres repository requires a valid gorm db instance")
	}

	if err := db.AutoMigrate(&replacemeRecord{}); err != nil {
		panic(err)
	}

	return &PostgresReplacemeRepository{db: db}
}

func (r *PostgresReplacemeRepository) Save(replaceme *models.Replaceme) error {
	if replaceme == nil {
		return errors.New("replaceme is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if replaceme.ID.Value() == "" ||
		replaceme.Title.Value() == "" ||
		replaceme.Content.Value() == "" ||
		replaceme.Poster.Value() == "" ||
		replaceme.Tags.Value() == nil ||
		replaceme.CreatedAt.IsZero() ||
		replaceme.UpdatedAt.IsZero() {
		return errors.New("replaceme is incomplete")
	}

	payload := replacemeRecord{
		ID:        replaceme.ID.Value(),
		Title:     replaceme.Title.Value(),
		Content:   replaceme.Content.Value(),
		Poster:    replaceme.Poster.Value(),
		Tags:      replaceme.Tags.Value(),
		CreatedAt: replaceme.CreatedAt,
		UpdatedAt: replaceme.UpdatedAt,
	}

	return r.db.Save(&payload).Error
}

func (r *PostgresReplacemeRepository) Get(id *types.ID) (*models.Replaceme, error) {
	if id == nil {
		return nil, errors.New("replaceme id is nil")
	}
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payload replacemeRecord
	result := r.db.Where("id = ?", id.Value()).First(&payload)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	replacemeID, err := types.NewIDFromString(payload.ID)
	if err != nil {
		return nil, err
	}

	replacemeTitle, err := types.NewTitle(payload.Title)
	if err != nil {
		return nil, err
	}

	replacemeContent, err := types.NewContent(payload.Content)
	if err != nil {
		return nil, err
	}

	replacemePoster, err := types.NewPoster(payload.Poster)
	if err != nil {
		return nil, err
	}

	replacemeTags, err := types.NewTags(payload.Tags)
	if err != nil {
		return nil, err
	}

	return &models.Replaceme{
		ID:        replacemeID,
		Title:     replacemeTitle,
		Content:   replacemeContent,
		Poster:    replacemePoster,
		Tags:      replacemeTags,
		CreatedAt: payload.CreatedAt,
		UpdatedAt: payload.UpdatedAt,
	}, nil
}

func (r *PostgresReplacemeRepository) Delete(id *types.ID) error {
	if id == nil {
		return errors.New("replaceme id is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	result := r.db.Delete(&replacemeRecord{}, "id = ?", id.Value())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresReplacemeRepository) Update(replaceme *models.Replaceme) error {
	if replaceme == nil {
		return errors.New("replaceme is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if replaceme.ID.Value() == "" ||
		replaceme.Title.Value() == "" ||
		replaceme.Content.Value() == "" ||
		replaceme.Poster.Value() == "" ||
		replaceme.Tags.Value() == nil ||
		replaceme.CreatedAt.IsZero() ||
		replaceme.UpdatedAt.IsZero() {
		return errors.New("replaceme is incomplete")
	}

	payload := replacemeRecord{
		ID:        replaceme.ID.Value(),
		Title:     replaceme.Title.Value(),
		Content:   replaceme.Content.Value(),
		Poster:    replaceme.Poster.Value(),
		Tags:      replaceme.Tags.Value(),
		CreatedAt: replaceme.CreatedAt,
		UpdatedAt: replaceme.UpdatedAt,
	}

	result := r.db.Model(&replacemeRecord{}).Where("id = ?", payload.ID).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresReplacemeRepository) List() ([]*models.Replaceme, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payloads []replacemeRecord
	if err := r.db.Find(&payloads).Error; err != nil {
		return nil, err
	}

	replacemes := make([]*models.Replaceme, 0, len(payloads))
	for _, payload := range payloads {
		replacemeID, err := types.NewIDFromString(payload.ID)
		if err != nil {
			return nil, err
		}

		replacemeTitle, err := types.NewTitle(payload.Title)
		if err != nil {
			return nil, err
		}

		replacemeContent, err := types.NewContent(payload.Content)
		if err != nil {
			return nil, err
		}

		replacemePoster, err := types.NewPoster(payload.Poster)
		if err != nil {
			return nil, err
		}

		replacemeTags, err := types.NewTags(payload.Tags)
		if err != nil {
			return nil, err
		}

		replacemes = append(replacemes, &models.Replaceme{
			ID:        replacemeID,
			Title:     replacemeTitle,
			Content:   replacemeContent,
			Poster:    replacemePoster,
			Tags:      replacemeTags,
			CreatedAt: payload.CreatedAt,
			UpdatedAt: payload.UpdatedAt,
		})
	}

	return replacemes, nil
}
