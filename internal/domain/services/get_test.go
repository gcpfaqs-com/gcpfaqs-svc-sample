package services_test

import (
	"gcpreplacemes-svc-replacemes/internal/domain/models"
	"gcpreplacemes-svc-replacemes/internal/domain/services"
	"gcpreplacemes-svc-replacemes/internal/domain/test"
	"gcpreplacemes-svc-replacemes/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAResource(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
	repository := test.NewStubRepository()

	sut := services.NewGetReplacemeUseCase(repository)
	resource, err := sut.Do(id)

	assert.Nil(t, err)
	assert.IsType(t, &models.Replaceme{}, resource)
	assert.True(t, repository.GetHaveBeenCalled(), sut)
}

func TestGetItShouldAnErrorOnInvalidId(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba"
	repository := test.NewStubRepository()

	sut := services.NewGetReplacemeUseCase(repository)
	_, err := sut.Do(id)

	assert.IsType(t, err, types.ErrInvalidID)
}
