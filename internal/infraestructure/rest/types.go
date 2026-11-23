package rest

import "time"

type replacemeDTO struct {
	ID        string    `json:"id,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Poster    string    `json:"poster"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
