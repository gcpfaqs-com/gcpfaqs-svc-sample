package types

import "errors"

var ErrInvalidPoster = errors.New("invalid Poster must be a non-empty string between 10 and 250 characters")

type Poster struct {
	value string
}

func NewPoster(s string) (Poster, error) {
	if len(s) < 10 || len(s) > 250 {
		return Poster{}, ErrInvalidPoster
	}

	return Poster{
		value: s,
	}, nil
}

func (c *Poster) Value() string {
	return c.value
}
