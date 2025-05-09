package git

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

var errGitNotFound = errors.New("git executable not found")

// Diff returns the staged diff via the git CLI.
// If git is missing or there are no staged changes, it falls back
// to a basic status summary from go-git.
func (s *Service) Diff() (string, error) {
	diff, err := s.execGitDiff("--staged")
	if err == nil && diff != "" {
		return diff, nil
	}
	if err != nil && !errors.Is(err, errGitNotFound) {
		return "", fmt.Errorf("exec git diff: %w", err)
	}

	// fallback to status summary
	return s.statusSummary()
}

func (s *Service) execGitDiff(args ...string) (string, error) {
	// whitelist only the --staged flag
	for _, arg := range args {
		if arg != "--staged" {
			return "", fmt.Errorf("invalid git diff argument: %s", arg)
		}
	}

	// ensure git is in PATH
	if _, err := exec.LookPath("git"); err != nil {
		return "", errGitNotFound
	}

	// launch subprocess with constant executable name
	cmd := exec.Command("git", append([]string{"diff"}, args...)...) //nolint:gosec
	cmd.Dir = s.path

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s failed: %w\n%s",
			strings.Join(args, " "), err, output)
	}
	return string(output), nil
}

// statusSummary builds a simple diff-like report based on go-git status.
func (s *Service) statusSummary() (string, error) {
	wt, err := s.repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("access worktree: %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return "", fmt.Errorf("retrieve status: %w", err)
	}

	var buf bytes.Buffer
	for path, st := range status {
		if st.Staging != git.Unmodified {
			stagingStatus := "Modified"
			switch st.Staging {
			case git.Untracked:
				stagingStatus = "Untracked"
			case git.Modified:
				stagingStatus = "Modified"
			case git.Added:
				stagingStatus = "Added"
			case git.Deleted:
				stagingStatus = "Deleted"
			case git.Renamed:
				stagingStatus = "Renamed"
			case git.Copied:
				stagingStatus = "Copied"
			case git.UpdatedButUnmerged:
				stagingStatus = "Conflicted"
			}
			_, err := fmt.Fprintf(&buf, "File: %s (Status: %s)\n", path, stagingStatus)
			if err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}
