package github

import (
	"context"

	"github.com/shurcooL/githubv4"
	"go.mattglei.ch/go.mattglei.ch/internal/secrets"
	"golang.org/x/oauth2"
)

func Client() (githubv4.Client, error) {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: secrets.ENV.GitHubToken})
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return *client, nil
}
