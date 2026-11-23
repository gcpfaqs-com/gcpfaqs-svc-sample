package services_test

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAList(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewListReplacemeUseCase(repository)
	list, err := sut.Do()

	assert.Nil(t, err)
	assert.IsType(t, []*models.Replaceme{}, list)
	assert.True(t, repository.ListHaveBeenCalled(), sut)
}
