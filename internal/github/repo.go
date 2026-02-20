package github

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
)

type Repository struct {
	Owner       string
	Name        string
	Description string
	Version     string
	Updated     time.Time
}

func FetchRepo(clients Clients, owner, name string) (Repository, error) {
	var q struct {
		Repository struct {
			Name        githubv4.String
			Description githubv4.String
			UpdatedAt   githubv4.DateTime

			Tags struct {
				Nodes []struct {
					Name githubv4.String
				}
			} `graphql:"refs(refPrefix: \"refs/tags/\", first: 1, orderBy: {field: TAG_COMMIT_DATE, direction: DESC})"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	vars := map[string]any{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	if err := clients.GraphQL.Query(context.Background(), &q, vars); err != nil {
		return Repository{}, fmt.Errorf("github graphql query for %s/%s: %w", owner, name, err)
	}

	version := ""
	if len(q.Repository.Tags.Nodes) > 0 && q.Repository.Tags.Nodes[0].Name != "" {
		version = string(q.Repository.Tags.Nodes[0].Name)
	}

	return Repository{
		Owner:       "gleich",
		Name:        string(q.Repository.Name),
		Description: string(q.Repository.Description),
		Version:     version,
		Updated:     q.Repository.UpdatedAt.Time,
	}, nil
}
