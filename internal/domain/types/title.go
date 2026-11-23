package types

import "errors"

var ErrInvalidTitle = errors.New("invalid title must be a non-empty string between 5 and 150 characters")

type Title struct {
	value string
}

func NewTitle(s string) (Title, error) {
	if len(s) < 5 || len(s) > 150 {
		return Title{}, ErrInvalidTitle
	}

	return Title{
		value: s,
	}, nil
}

func (t *Title) Value() string {
	return t.value
}
