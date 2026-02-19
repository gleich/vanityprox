package main

import (
	"time"

	"go.mattglei.ch/timber"
	"go.mattglei.ch/vanityprox/internal/api"
	"go.mattglei.ch/vanityprox/internal/conf"
	"go.mattglei.ch/vanityprox/internal/github"
	"go.mattglei.ch/vanityprox/internal/pkg"
	"go.mattglei.ch/vanityprox/internal/secrets"
)

func main() {
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		timber.Fatal(err, "failed to read new york time location")
	}
	timber.Timezone(ny)
	timber.TimeFormat("01/02 03:04:05pm MST")

	config, err := conf.Read()
	if err != nil {
		timber.Fatal(err, "failed to load configuration")
	}

	config.Log()

	secrets.Load()

	githubClient, err := github.Client()
	if err != nil {
		timber.Fatal(err, "failed to create github client")
	}

	packages, err := pkg.Setup(config, &githubClient)
	if err != nil {
		timber.Fatal(err, "failed to setup packages")
	}

	server := api.Setup(config, packages)

	timber.Donef("starting server on 0.0.0.0%s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		timber.Fatal(err, "failed to start server")
	}
}
