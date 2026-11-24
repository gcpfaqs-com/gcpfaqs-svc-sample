package services_test

import (
	"time"

	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newTitle string = "any valid title"
var newContent string = "any valid content"
var newPoster string = "any valid poster"
var newTags []string = []string{"atag"}

func TestItShouldSaveAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(newTitle, newContent, newPoster, newTags, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.SaveHaveBeenCalled(), sut)
}

func TestItShouldResultAnErrorOnInvalidTitle(t *testing.T) {
	newTitle = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(newTitle, newContent, newPoster, newTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTitle, err)
}

func TestItShouldResultAnErrorOnInvalidContent(t *testing.T) {
	newContent = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(newTitle, newContent, newPoster, newTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidContent, err)
}

func TestItShouldResultAnErrorOnInvalidPoster(t *testing.T) {
	newPoster = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(newTitle, newContent, newPoster, newTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidPoster, err)
}
func TestItShouldResultAnErrorOnInvalidTags(t *testing.T) {
	newTags = []string{}
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(newTitle, newContent, newPoster, newTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTags, err)
}
