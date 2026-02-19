package github

import (
	"context"

	githubREST "github.com/google/go-github/v83/github"
	"github.com/shurcooL/githubv4"
	"go.mattglei.ch/go.mattglei.ch/internal/secrets"
	"golang.org/x/oauth2"
)

type Clients struct {
	REST    *githubREST.Client
	GraphQL *githubv4.Client
}

func CreateClients() (Clients, error) {
	var (
		token      = secrets.ENV.GitHubToken
		src        = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(context.Background(), src)
	)

	graphql := githubv4.NewClient(httpClient)
	rest := githubREST.NewClient(httpClient)
	return Clients{
		REST:    rest,
		GraphQL: graphql,
	}, nil
}
