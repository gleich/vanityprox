package pkg

import (
	"fmt"
	"time"

	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/github"
	"go.mattglei.ch/timber"
)

func Setup(
	config conf.Config,
	clients github.Clients,
) (*Packages, error) {
	start := time.Now()
	p := Packages{}

	err := github.SetupCloneFolder()
	if err != nil {
		timber.Fatal(err, "failed to setup folder for cloning")
	}

	for _, name := range config.Packages {
		repo, err := github.FetchRepo(clients, "gleich", name)
		if err != nil {
			return &Packages{}, err
		}

		err = repo.Subscribe(clients)
		if err != nil {
			return &Packages{}, fmt.Errorf("creating webhook: %w", err)
		}

		err = repo.Clone()
		if err != nil {
			return &Packages{}, fmt.Errorf("cloning repo: %w", err)
		}

		p.packages = append(p.packages, repo)
	}
	timber.DoneSince(start, "loaded info for", len(p.packages), "packages")
	return &p, nil
}
