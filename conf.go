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
	Host          string `env:"HOST,required"`
	SourcePrefix  string `env:"SOURCE_PREFIX,required"`
	Favicon       string `env:"FAVICON"`
	RootRedirect  string `env:"ROOT_REDIRECT"`
	LogTimezone   string `env:"LOG_TIMEZONE"`
	LogTimeFormat string `env:"LOG_TIME_FORMAT"`
}

func readConfig() (config, error) {
	if _, err := os.Stat(".env"); !errors.Is(err, fs.ErrNotExist) {
		err = godotenv.Load()
		if err != nil {
			return config{}, fmt.Errorf("reading .env file: %w", err)
		}
	}

	conf, err := env.ParseAsWithOptions[config](env.Options{Prefix: "VANITYPROX_"})
	if err != nil {
		return config{}, fmt.Errorf("parsing config from environment variables: %w", err)
	}

	// ensure that source prefix is formatted properly
	sourceURL, err := url.Parse(conf.SourcePrefix)
	if err != nil {
		return config{}, fmt.Errorf("parsing source prefix URL: %w", err)
	}
	sourcePrefix, err := url.JoinPath(sourceURL.Host, sourceURL.Path)
	if err != nil {
		return config{}, fmt.Errorf("creating source prefix from URL: %w", err)
	}
	conf.SourcePrefix = sourcePrefix

	hostURL, err := url.Parse(conf.Host)
	if err != nil {
		return config{}, fmt.Errorf("%w failed to parse host", err)
	}
	conf.Host = hostURL.Host

	return conf, nil
}

func (c config) log() {
	timber.Info("           host =", c.Host)
	timber.Info("  source prefix =", c.SourcePrefix)
	if c.Favicon != "" {
		timber.Info("        favicon =", c.Favicon)
	}
	if c.RootRedirect != "" {
		timber.Info("  root redirect =", c.RootRedirect)
	}
	if c.LogTimezone != "" {
		timber.Info("   log timezone =", c.LogTimezone)
	}
	if c.LogTimeFormat != "" {
		timber.Info("log time format =", c.LogTimeFormat)
	}
}
