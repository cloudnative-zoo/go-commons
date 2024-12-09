package main

import (
	"context"
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/github"
)

func main() {
	ctx := context.Background()
	githbSvc, err := github.New(
		github.WithToken(""), // export GITHUB_TOKEN=your_token
	)
	if err != nil {
		panic(err)
	}
	rateLimits, err := githbSvc.CheckRateLimit(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to check rate limits")
	}
	slog.With("rateLimits", rateLimits).Info("rate limits checked successfully")
	myOrgs, err := githbSvc.ListOrganizations(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to list organizations")
	}
	slog.With("myOrgs", myOrgs).Info("organizations listed successfully")
}
