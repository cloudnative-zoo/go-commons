package main

import (
	"log/slog"

	"github.com/cloudnative-zoo/go-commons/gitlab"
)

func main() {
	svc, err := gitlab.New(
		gitlab.WithToken(""), // export GITLAB_TOKEN=your_token
	)
	if err != nil {
		panic(err)
	}

	userId, err := svc.GetUserID()
	if err != nil {
		slog.Error("failed to get user id: %w", err)
	}
	slog.Info("user id: %d", userId)

	projects, err := svc.ListOwnedProjects()
	if err != nil {
		slog.Error("failed to list projects: %w", err)
	}

	for _, project := range projects {
		slog.With("Name", project.Name, "ID", project.ID, "SShUrl", project.SSHURLToRepo).Info("project")
	}

}
