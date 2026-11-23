package types

import (
	"errors"
)

var ErrInvalidVideoPlaylistID = errors.New("invalid video playlist ID must be an UUID v4 string")

type VideoPlaylistID struct {
	value string
}

func NewVideoPlaylistID(s string) (VideoPlaylistID, error) {
	id, err := NewIDFromString(s)
	if err != nil {
		return VideoPlaylistID{}, ErrInvalidVideoPlaylistID
	}

	return VideoPlaylistID{value: id.Value()}, nil
}

func (v *VideoPlaylistID) Value() string {
	return v.value
}

func (v *VideoPlaylistID) IsValid() bool {
	_, err := NewIDFromString(v.value)
	return err == nil
}
