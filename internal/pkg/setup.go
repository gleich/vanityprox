package pkg

import (
	"time"

	"github.com/shurcooL/githubv4"
	"go.mattglei.ch/timber"
	"go.mattglei.ch/vanityprox/internal/conf"
	"go.mattglei.ch/vanityprox/internal/github"
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
