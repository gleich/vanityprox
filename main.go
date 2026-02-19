package main

import (
	"time"

	"go.mattglei.ch/go.mattglei.ch/internal/api"
	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/github"
	"go.mattglei.ch/go.mattglei.ch/internal/pkg"
	"go.mattglei.ch/go.mattglei.ch/internal/secrets"
	"go.mattglei.ch/timber"
)

func main() {
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		timber.Fatal(err, "failed to read new york time location")
	}
	timber.Timezone(ny)
	timber.TimeFormat("01/02 03:04:05 PM  MST")

	config, err := conf.Read()
	if err != nil {
		timber.Fatal(err, "failed to load configuration")
	}

	config.Log()

	secrets.Load()

	clients, err := github.CreateClients()
	if err != nil {
		timber.Fatal(err, "failed to create github client")
	}

	packages, err := pkg.Setup(config, clients)
	if err != nil {
		timber.Fatal(err, "failed to setup packages")
	}

	server := api.Setup(config, clients, packages)

	timber.Donef("starting server on 0.0.0.0%s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		timber.Fatal(err, "failed to start server")
	}
}
