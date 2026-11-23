package types

import (
	"errors"
)

var ErrInvalidYoutubeID = errors.New("invalid youtube ID must be a string")

type YoutubeID struct {
	value string
}

func NewYoutubeID(youtubeID string) (YoutubeID, error) {
	if youtubeID == "" {
		return YoutubeID{}, ErrInvalidYoutubeID
	}

	return YoutubeID{value: youtubeID}, nil
}

func (v *YoutubeID) Value() string {
	return v.value
}
