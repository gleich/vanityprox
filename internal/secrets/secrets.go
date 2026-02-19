package secrets

import (
	"errors"
	"io/fs"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.mattglei.ch/timber"
)

var ENV Secrets

type Secrets struct {
	GitHubToken         string `env:"GITHUB_TOKEN"`
	GitHubWebhookSecret string `env:"GITHUB_WEBHOOK_SECRET"`
}

func Load() {
	start := time.Now()
	err := godotenv.Load()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		timber.Fatal(err, "loading .env file failed")
	}

	secrets, err := env.ParseAsWithOptions[Secrets](
		env.Options{RequiredIfNoDef: true},
	)
	if err != nil {
		timber.Fatal(err, "parsing required env vars failed")
	}
	ENV = secrets
	timber.DoneSince(start, "loaded secrets")
}
