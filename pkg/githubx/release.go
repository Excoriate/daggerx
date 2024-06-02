package githubx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// FetchLatestRelease fetches the latest release from the GitHub repository.
//
// Returns:
//   - A string representing the latest release tag.
//   - An error if the latest release cannot be fetched.
func (gh *GHClient) FetchLatestRelease() (string, error) {
	ctx := context.Background()

	var tc *http.Client
	if gh.cfg.GetToken() != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: gh.cfg.GetToken()},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(tc)

	release, _, err := client.Repositories.GetLatestRelease(ctx, gh.cfg.GetOwner(), gh.cfg.GetRepo())
	if err != nil {
		return "", fmt.Errorf("failed to fetch the latest release: %w", err)
	}

	return release.GetTagName(), nil
}
