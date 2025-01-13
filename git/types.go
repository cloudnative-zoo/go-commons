package git

import (
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Service struct {
	auth            transport.AuthMethod
	repo            *gogit.Repository
	url             string
	path            string
	branch          string
	cloneIfNotExist bool
	progress        sideband.Progress
}
