// Package githubx provides utilities for interacting with GitHub's API,
// specifically for fetching information about releases in a GitHub repository.
//
// This package includes functionality to authenticate with GitHub using an
// OAuth2 token, and to retrieve the latest release information from a specified
// repository. It leverages the go-github library to interact with GitHub's API.
//
// Example usage:
//
//	import (
//	    "fmt"
//	    "github.com/yourusername/yourrepo/pkg/githubx"
//	)
//
//	func main() {
//	    client := githubx.NewGHClient("your-github-token", "owner", "repo")
//	    latestRelease, err := client.FetchLatestRelease()
//	    if err != nil {
//	        fmt.Println("Error fetching latest release:", err)
//	        return
//	    }
//	    fmt.Println("Latest release:", latestRelease)
//	}
//
// The above example demonstrates how to create a new GHClient instance and
// fetch the latest release tag from a specified GitHub repository.
//
// Note: Ensure that you have a valid GitHub token with the necessary
// permissions to access the repository's release information.
//
// For more details on GitHub's API, refer to the official documentation:
// https://docs.github.com/en/rest/reference/repos#releases
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
