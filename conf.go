package main

import (
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.mattglei.ch/timber"
)

type config struct {
	Host          string `env:"HOST,required"`          // required
	SourcePrefix  string `env:"SOURCE_PREFIX,required"` // required
	Favicon       string `env:"FAVICON"`                // optional
	RootRedirect  string `env:"ROOT_REDIRECT"`          // optional
	LogTimezone   string `env:"LOG_TIMEZONE"`           // optional
	LogTimeFormat string `env:"LOG_TIME_FORMAT"`        // optional
}

func readConfig() (config, error) {
	if _, err := os.Stat(".env"); !errors.Is(err, fs.ErrNotExist) {
		err = godotenv.Load()
		if err != nil {
			return config{}, fmt.Errorf("%w failed to read from .env file", err)
		}
	}

	conf, err := env.ParseAsWithOptions[config](env.Options{Prefix: "VANITYPROX_"})
	if err != nil {
		return config{}, fmt.Errorf("%w failed to parse config from environment variables", err)
	}

	// ensure that source prefix is formatted properly
	sourceURL, err := url.Parse(conf.SourcePrefix)
	if err != nil {
		return config{}, fmt.Errorf("%w failed to parse source prefix URL", err)
	}
	sourcePrefix, err := url.JoinPath(sourceURL.Host, sourceURL.Path)
	if err != nil {
		return config{}, fmt.Errorf("%w failed to create source prefix from URL", err)
	}
	conf.SourcePrefix = sourcePrefix

	hostURL, err := url.Parse(conf.Host)
	if err != nil {
		return config{}, fmt.Errorf("%w failed to parse host", err)
	}
	conf.Host = hostURL.Host

	return conf, nil
}

func logConfig(conf config) {
	timber.Info("           host =", conf.Host)
	timber.Info("  source prefix =", conf.SourcePrefix)
	if conf.Favicon != "" {
		timber.Info("        favicon =", conf.Favicon)
	}
	if conf.RootRedirect != "" {
		timber.Info("  root redirect =", conf.RootRedirect)
	}
	if conf.LogTimezone != "" {
		timber.Info("   log timezone =", conf.LogTimezone)
	}
	if conf.LogTimeFormat != "" {
		timber.Info("log time format =", conf.LogTimeFormat)
	}
}
