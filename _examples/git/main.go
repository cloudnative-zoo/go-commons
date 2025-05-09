package main

import (
	"context"
	"log/slog"
	"os"
	"path"

	"github.com/cloudnative-zoo/go-commons/git"
)

func main() {
	ctx := context.Background()
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		slog.With("error", err).Error("failed to get user home directory")
		os.Exit(1)
	}
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		slog.With("error", err).Error("failed to get current working directory")
		os.Exit(1)
	}
	gitSvc, err := git.New(
		ctx,
		git.WithSSHKeyPath(path.Join(userHomeDir, ".ssh", "github_hassnatahmad"), ""),
		git.WithRepoPath(currentWorkingDir),
		// git.WithURL(fmt.Sprintf("git@github.com:%s/%s.git", githubOrg, githubRepo)),
	)
	if err != nil {
		slog.With("error", err).Error("failed to create git service")
		return
	}
	// Pull the repository.
	// pull(ctx, gitSvc)
	// Check the status of the repository.
	status(ctx, gitSvc)
	// List branches.
	listBranches(ctx, gitSvc)
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

func status(_ context.Context, gitSvc *git.Service) {
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

func listBranches(_ context.Context, gitSvc *git.Service) {
	localBranches, err := gitSvc.ListLocalBranches()
	if err != nil {
		slog.With("error", err).Error("failed to list local branches")
	}
	remoteBranches, err := gitSvc.ListRemoteBranches()
	if err != nil {
		slog.With("error", err).Error("failed to list remote branches")
	}
	slog.With("LocalBranches", localBranches).With("RemoteBranches", remoteBranches).Info("branches")
}
