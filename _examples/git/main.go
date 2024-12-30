package main

import (
	"context"
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/git"
)

func main() {
	ctx := context.Background()
	gitSvc, err := git.New(
		git.WithSSHKeyPath("/Users/hassnat/.ssh/hassnatahmad_ssh", ""),
		git.WithRepoPath("/Users/hassnat/development/github/cloudnative-zoo/go-commons"),
		git.WithURL("git@github.com:cloudnative-zoo/go-commons.git"),
	)
	if err != nil {
		slog.With("error", err).Error("failed to create git service")
		return
	}
	status(ctx, gitSvc)
	diff(ctx, gitSvc)
}

func cloneOrPull(ctx context.Context, gitSvc *git.Service) {
	err := gitSvc.CloneOrPull(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to clone repository")
	}
	slog.Info("repository cloned successfully")
}

func diff(ctx context.Context, gitSvc *git.Service) {
	diff, err := gitSvc.Diff(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to get git diff")
	}
	slog.With("diff", diff).Info("git diff")
}

func fetch(ctx context.Context, gitSvc *git.Service) {
	err := gitSvc.Fetch(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to fetch repository")
	}
	slog.Info("repository fetched successfully")
}

func status(ctx context.Context, gitSvc *git.Service) {
	result, err := gitSvc.Status()
	if err != nil {
		slog.With("error", err).Error("failed to get git status")
	}
	slog.With("Added", result.Added).With("Modified", result.Modified).With("Deleted", result.Deleted).Info("git status")
}
