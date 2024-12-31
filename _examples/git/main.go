package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/cloudnative-zoo/go-commons/git"
)

func main() {
	ctx := context.Background()
	homeDir := os.Getenv("HOME")
	githubOrg := "cloudnative-zoo"
	githubRepo := "go-commons"
	repoPath := path.Join(homeDir, "development", githubOrg, githubRepo)
	gitSvc, err := git.New(
		ctx,
		git.WithSSHKeyPath(path.Join(homeDir, ".ssh", "github_hassnatahmad"), ""),
		git.WithRepoPath(repoPath),
		git.WithURL(fmt.Sprintf("git@github.com:%s/%s.git", githubOrg, githubRepo)),
	)
	if err != nil {
		slog.With("error", err).Error("failed to create git service")
		return
	}
	// Pull the repository.
	pull(ctx, gitSvc)
	// Check the status of the repository.
	status(ctx, gitSvc)
}

func pull(ctx context.Context, gitSvc *git.Service) {
	err := gitSvc.Pull(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to clone repository")
	}
	slog.Info("repository pulled successfully")
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
	if result == nil {
		slog.Info("repository is clean")
		return
	}
	slog.With("Added", result.Added).With("Modified", result.Modified).With("Deleted", result.Deleted).Info("git status")
}
