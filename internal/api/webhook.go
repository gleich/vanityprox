package api

import (
	"fmt"
	"net/http"

	githubREST "github.com/google/go-github/v83/github"
	"go.mattglei.ch/go.mattglei.ch/internal/github"
	"go.mattglei.ch/go.mattglei.ch/internal/pkg"
	"go.mattglei.ch/go.mattglei.ch/internal/secrets"
	"go.mattglei.ch/go.mattglei.ch/internal/util"
	"go.mattglei.ch/timber"
)

func webhookEndpoint(
	w http.ResponseWriter,
	r *http.Request,
	clients github.Clients,
	packages *pkg.Packages,
) {
	payload, err := githubREST.ValidatePayload(r, []byte(secrets.ENV.GitHubWebhookSecret))
	if err != nil {
		http.Error(w, "invalid signature", http.StatusUnauthorized)
	}

	event, err := githubREST.ParseWebHook(githubREST.WebHookType(r), payload)
	if err != nil {
		util.InternalServerError(w, err)
		return
	}

	var (
		name  string
		owner string
	)
	switch e := event.(type) {
	case *githubREST.PushEvent:
		repo := e.GetRepo()
		name = *repo.Name
		owner = *repo.GetOwner().Name
	case *githubREST.ReleaseEvent:
		repo := e.GetRepo()
		name = *repo.Name
		owner = *repo.GetOwner().Name
	case *githubREST.RepositoryEvent:
		repo := e.GetRepo()
		name = *repo.Name
		owner = *repo.GetOwner().Name
	}

	repo := packages.Get(name)
	if repo != nil {
		err = repo.Update(clients)
		if err != nil {
			util.InternalServerError(w, err)
			return
		}
		timber.Done("updated", name)
	} else if name != "" && owner != "" {

		err = github.Unsubscribe(clients, owner, name)
		if err != nil {
			util.InternalServerError(w, fmt.Errorf("unsubscribing from %s/%s: %w", owner, name, err))
		}
		timber.Done("removed webhook for", name)
	}
}
