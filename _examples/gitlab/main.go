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
		slog.Error("failed to get user id", slog.Any("error", err))
	}
	slog.With("UserID", userId).Info("user id")

	projects, err := svc.ListOwnedProjects()
	if err != nil {
		slog.Error("failed to list projects", slog.Any("error", err))
	}

	for _, project := range projects {
		slog.With("Name", project.Name, "ID", project.ID, "SShUrl", project.SSHURLToRepo).Info("project")
	}

}
