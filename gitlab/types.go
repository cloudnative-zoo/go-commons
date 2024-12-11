package gitlab

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type Service struct {
	client             *gitlab.Client
	paginationMaxLimit int
}
