package services_test

import (
	"time"

	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldSaveAResource(t *testing.T) {
	title := "any valid title"
	content := "any valid content"
	poster := "any valid poster"
	tags := []string{"atag"}
	repository := test.NewStubRepository()

	sut := services.NewSaveReplacemeUseCase(repository)
	_, err := sut.Do(title, content, poster, tags, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.SaveHaveBeenCalled(), sut)
}
