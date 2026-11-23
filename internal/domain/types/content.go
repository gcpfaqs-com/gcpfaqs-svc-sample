package types

import "errors"

var ErrInvalidContent = errors.New("invalid content must be a non-empty string between 10 and 10000 characters")

type Content struct {
	value string
}

func NewContent(s string) (Content, error) {
	if len(s) < 10 || len(s) > 10000 {
		return Content{}, ErrInvalidContent
	}

	return Content{
		value: s,
	}, nil
}

func (c *Content) Value() string {
	return c.value
}
