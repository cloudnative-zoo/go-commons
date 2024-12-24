package main

import (
	"context"
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/github"
)

func main() {
	ctx := context.Background()
	svc, err := github.New(
		github.WithToken(""), // export GITHUB_TOKEN=your_token
	)
	if err != nil {
		panic(err)
	}
	// GetRateLimit(ctx,svc)
	// ListOrganizations(ctx,svc)
	ListOpenPullRequestsAndMerge(ctx, svc, "cloudnative-zoo", "argocd-gitops-addons")
}

func GetRateLimit(ctx context.Context, svc *github.Service) {
	rateLimits, err := svc.CheckRateLimit(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to check rate limits")
	}
	slog.With("rateLimits", rateLimits).Info("rate limits checked successfully")
}

func ListOrganizations(ctx context.Context, svc *github.Service) {
	myOrgs, err := svc.ListOrganizations(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to list organizations")
	}
	slog.With("myOrgs", myOrgs).Info("organizations listed successfully")
}

func ListOpenPullRequestsAndMerge(ctx context.Context, svc *github.Service, owner, repo string) {
	prs, err := svc.ListOpenPullRequests(ctx, owner, repo)
	if err != nil {
		slog.With("error", err).Error("failed to list pull requests")
	}
	// slog.With("prs", prs).Info("pull requests listed successfully")
	for _, pr := range prs {
		// GetPullRequest(ctx,svc,owner,repo,pr.GetNumber())
		result, err := svc.MergePullRequest(ctx, owner, repo, pr.GetNumber())
		if err != nil {
			slog.With("error", err).Error("failed to merge pull request")
			return
		}
		slog.With("PR#", pr.GetNumber(), "merged", result.GetMerged(), "message", result.GetMessage()).Info("pull request merge result")
	}
}
