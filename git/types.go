package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Service struct {
	auth     transport.AuthMethod
	repo     *git.Repository
	url      string
	path     string
	branch   string
	progress sideband.Progress
}
