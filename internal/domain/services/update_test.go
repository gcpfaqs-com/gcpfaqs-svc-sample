package services_test

import (
	"time"

	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var updateId string = "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
var updateTitle string = "any valid title"
var updateContent string = "any valid content"
var updatePoster string = "any valid poster"
var updateTags []string = []string{"atag"}

func TestItShouldUpdateAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(updateId, updateTitle, updateContent, updatePoster, updateTags, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalled(), sut)
}

func TestUpdateItShouldResultAnErrorOnInvalidTitle(t *testing.T) {
	title := ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(updateId, title, updateContent, updatePoster, updateTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTitle, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidContent(t *testing.T) {
	content := ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(updateId, updateTitle, content, updatePoster, updateTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidContent, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidPoster(t *testing.T) {
	poster := ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(updateId, updateTitle, updateContent, poster, updateTags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidPoster, err)
}
func TestUpdateItShouldResultAnErrorOnInvalidTags(t *testing.T) {
	tags := []string{}
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(updateId, updateTitle, updateContent, updatePoster, tags, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTags, err)
}
