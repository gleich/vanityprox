package pkg

import (
	"time"

	"github.com/shurcooL/githubv4"
	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/github"
	"go.mattglei.ch/timber"
)

func Setup(config conf.Config, client *githubv4.Client) (*Packages, error) {
	start := time.Now()
	p := Packages{}
	for _, name := range config.Packages {
		repo, err := github.FetchRepo(client, "gleich", name)
		if err != nil {
			return &Packages{}, err
		}
		p.packages = append(p.packages, repo)
	}
	timber.DoneSince(start, "loaded info for", len(p.packages), "packages")
	return &p, nil
}
