package models

import (
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"time"
)

type Replaceme struct {
	ID        types.ID      `json:"id"`
	Title     types.Title   `json:"title"`
	Content   types.Content `json:"content"`
	Poster    types.Poster  `json:"poster"`
	Tags      types.Tags    `json:"tags"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func NewReplaceme(
	title string,
	content string,
	poster string,
	tags []string,
	createdAt time.Time,
	updatedAt time.Time) (*Replaceme, error) {
	replacemeTitle, err := types.NewTitle(title)
	if err != nil {
		return nil, err
	}

	replacemeContent, err := types.NewContent(content)
	if err != nil {
		return nil, err
	}

	replacemePoster, err := types.NewPoster(poster)
	if err != nil {
		return nil, err
	}

	replacemeTags, err := types.NewTags(tags)
	if err != nil {
		return nil, err
	}

	replacemeID, err := types.NewID()
	if err != nil {
		return nil, err
	}

	return &Replaceme{
		ID:        replacemeID,
		Title:     replacemeTitle,
		Content:   replacemeContent,
		Poster:    replacemePoster,
		Tags:      replacemeTags,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
