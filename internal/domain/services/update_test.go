package services_test

import (
	"time"

	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldUpdateAResource(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
	title := "any valid title"
	content := "any valid content"
	poster := "any valid poster"
	tags := []string{"atag"}
	repository := test.NewStubRepository()

	sut := services.NewUpdateReplacemeUseCase(repository)
	err := sut.Do(id, title, content, poster, tags, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalled(), sut)
}
