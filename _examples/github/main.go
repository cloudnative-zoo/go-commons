package main

import (
	"context"
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/github"
)

func main() {
	ctx := context.Background()
	githubService, err := github.New(
		github.WithToken(""), // export GITHUB_TOKEN=your_token
	)
	if err != nil {
		panic(err)
	}
	// GetRateLimit(ctx,githubService)
	// ListOrganizations(ctx,githubService)
	ListOpenPullRequestsAndMerge(ctx, githubService, "cloudnative-zoo", "argocd-gitops-addons")
}

func GetRateLimit(ctx context.Context, githubService *github.Service) {
	rateLimits, err := githubService.CheckRateLimit(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to check rate limits")
	}
	slog.With("rateLimits", rateLimits).Info("rate limits checked successfully")
}

func ListOrganizations(ctx context.Context, githubService *github.Service) {
	myOrgs, err := githubService.ListOrganizations(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to list organizations")
	}
	slog.With("myOrgs", myOrgs).Info("organizations listed successfully")
}

func ListOpenPullRequestsAndMerge(ctx context.Context, githubService *github.Service, owner, repo string) {
	prs, err := githubService.ListOpenPullRequests(ctx, owner, repo)
	if err != nil {
		slog.With("error", err).Error("failed to list pull requests")
	}
	// slog.With("prs", prs).Info("pull requests listed successfully")
	for _, pr := range prs {
		// GetPullRequest(ctx,githubService,owner,repo,pr.GetNumber())
		result, err := githubService.MergePullRequest(ctx, owner, repo, pr.GetNumber(), "")
		if err != nil {
			slog.With("error", err, "PR#", pr.GetNumber()).Error("failed to merge pull request")
		}
		slog.With("PR#", pr.GetNumber(), "merged", result.GetMerged(), "message", result.GetMessage()).Info("pull request merge result")
	}
}
