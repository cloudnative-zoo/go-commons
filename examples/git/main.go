package main

import (
	"context"
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/git"
)

func main() {
	ctx := context.Background()
	gitSvc, err := git.New(
		git.WithSSHKeyPath("/Users/hassnat/.ssh/edu_api_git", ""),
		git.WithRepoPath("/Users/hassnat/development/test/edyouapp-api"),
		git.WithURL("git@github.com:o2-edyou/edyouapp-api.git"),
	)
	if err != nil {
		panic(err)
	}
	err = gitSvc.CloneOrPull(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to clone repository")
	}
	slog.Info("repository cloned successfully")
}
