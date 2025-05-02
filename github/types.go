package github

import (
	"github.com/google/go-github/v71/github"
)

type Service struct {
	client             *github.Client
	paginationMaxLimit int
}
