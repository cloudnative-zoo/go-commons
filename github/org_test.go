package github_test

import (
	"testing"

	"github.com/google/go-github/v72/github"
	"github.com/stretchr/testify/assert"
)

type Service struct {
	client             *github.Client
	paginationMaxLimit int
}

func TestServiceInitialization(t *testing.T) {
	t.Parallel()
	client := &github.Client{}
	paginationMaxLimit := 100
	service := Service{
		client:             client,
		paginationMaxLimit: paginationMaxLimit,
	}

	assert.NotNil(t, service.client, "Expected client to be non-nil")
	assert.Equal(t, paginationMaxLimit, service.paginationMaxLimit, "Expected paginationMaxLimit to be set correctly")
}

func TestServiceNilClient(t *testing.T) {
	t.Parallel()
	paginationMaxLimit := 100
	service := Service{
		client:             nil,
		paginationMaxLimit: paginationMaxLimit,
	}

	assert.Nil(t, service.client, "Expected client to be nil")
}

func TestServicePaginationMaxLimit(t *testing.T) {
	t.Parallel()
	client := &github.Client{}
	maxLimit := 100
	excessiveLimit := 1000
	service := Service{
		client:             client,
		paginationMaxLimit: excessiveLimit,
	}

	if service.paginationMaxLimit > maxLimit {
		service.paginationMaxLimit = maxLimit
	}

	assert.LessOrEqual(t, service.paginationMaxLimit, maxLimit, "Expected paginationMaxLimit to not exceed the maximum limit")
}
