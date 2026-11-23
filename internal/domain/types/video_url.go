package types

import (
	"errors"
	"strings"
)

var ErrInvalidVideoURL = errors.New("invalid video URL must be a non-empty string")

type VideoURL struct {
	value string
}

func NewVideoURL(s string) (VideoURL, error) {
	if strings.TrimSpace(s) == "" {
		return VideoURL{}, ErrInvalidVideoURL
	}

	return VideoURL{value: s}, nil
}

func (v *VideoURL) Value() string {
	return v.value
}
