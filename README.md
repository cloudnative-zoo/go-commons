# go-commons

`go-commons` is a versatile collection of utility packages that simplify common tasks related to Git and GitHub. 

It provides high-level abstractions for Git operations, GitHub API interactions, and general development utilities like directory management and environment variable handling.

---

## Features

### **GitHub API Integration**
Easily interact with GitHub using the `go-github` library. Key features include:
- Listing user organizations and repositories.
- Checking rate limits.
- Managing authentication via tokens or SSH keys.

### **Git Operations**
Streamline Git-related tasks with `go-git`, including:
- Cloning repositories.
- Updating local repositories.
- Supporting both HTTPS and SSH protocols for authentication.

### **Utility Functions**
- Directory management: Check for existence and create directories if necessary.
- Environment variable handling: Fetch variables with default fallback options.

---

## Installation

To add `go-commons` to your project, run:

```bash
go get -u github.com/cloudnative-zoo/go-commons
```

---

## Usage

### **GitHub API Service Example**

```go
package main

import (
    "context"
    "log/slog"

    "github.com/cloudnative-zoo/go-commons/github"
)

func main() {
    ctx := context.Background()

    githubService, err := github.New(
        github.WithToken("your_github_token"),
    )
    if err != nil {
        slog.With("error", err).Error("Failed to initialize GitHub service")
        return
    }

    orgs, err := githubService.ListOrganizations(ctx)
    if err != nil {
        slog.With("error", err).Error("Failed to list organizations")
        return
    }

    slog.With("organizations", orgs).Info("Retrieved organizations")
}
```

### **Git Service Example**

```go
package main

import (
    "context"
    "log/slog"

    "github.com/cloudnative-zoo/go-commons/git"
)

func main() {
    ctx := context.Background()

    gitService, err := git.New(
        git.WithSSHKeyPath("/path/to/ssh/key", ""),
        git.WithRepoPath("/path/to/local/repo"),
        git.WithURL("git@github.com:user/repo.git"),
    )
    if err != nil {
        slog.With("error", err).Error("Failed to initialize Git service")
        return
    }

    err = gitService.CloneOrPull(ctx)
    if err != nil {
        slog.With("error", err).Error("Failed to clone or pull repository")
        return
    }

    slog.Info("Repository synchronized successfully")
}
```

---

## Development

### **Prerequisites**
- Install [Go](https://go.dev/)
- Install [pre-commit](https://pre-commit.com/)

### **Configuration**

#### **Taskfile**
A `Taskfile.yaml` is included to streamline common development tasks:
- Pre-commit checks
- Updating Go modules
- Formatting code
- Running `go vet`
- Running `golangci-lint`

#### **Pre-commit Hooks**
The `.pre-commit-config.yaml` file configures hooks for:
- **Sensitive data detection**: `gitleaks`
- **Formatting tools**: End-of-file fixer, trailing whitespace remover, JSON formatter
- **Linting tools**: `golangci-lint`, `gofmt`, `goimports`
- **Spelling and syntax checks**

---

## Contributing

We welcome contributions! Follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-branch
   ```
3. Make your changes.
4. Commit your changes:
   ```bash
   git commit -am 'Add new feature'
   ```
5. Push to the branch:
   ```bash
   git push origin feature-branch
   ```
6. Open a pull request.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgements

We are grateful for the following libraries and tools that power this project:
- [Cobra](https://github.com/spf13/cobra) for CLI commands
- [go-git](https://github.com/go-git/go-git) for Git operations
- [go-github](https://github.com/google/go-github) for GitHub API interactions
- [pre-commit](https://pre-commit.com/) for managing pre-commit hooks

---

## Contact

For questions, suggestions, or feedback, please [open an issue](https://github.com/cloudnative-zoo/go-commons/issues) on the GitHub repository.

---

## Authors

- **Hassnat Ahmad** - [github](https://github.com/hassnatahmad)