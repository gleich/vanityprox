package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pelletier/go-toml/v2"
	"go.mattglei.ch/timber"
)

type config struct {
	Host         string `toml:"host"`
	SourcePrefix string `toml:"source_prefix"`
	Favicon      string `toml:"favicon"`
	RootRedirect string `toml:"root_redirect"`
	Logs         struct {
		Timezone   string `toml:"timezone"`
		TimeFormat string `toml:"time_format"`
	} `toml:"logs"`
	Packages []string `toml:"packages"`
}

func readConfig() (config, error) {
	filename := "vanityprox.toml"
	bin, err := os.ReadFile(filename)
	if err != nil {
		return config{}, fmt.Errorf("reading from %s: %w", filename, err)
	}

	var conf config
	err = toml.Unmarshal(bin, &conf)
	if err != nil {
		return config{}, fmt.Errorf("unmarshaling toml: %w", err)
	}

	// ensure that source prefix is formatted properly
	sourceURL, err := url.Parse(conf.SourcePrefix)
	if err != nil {
		return config{}, fmt.Errorf("parsing source prefix URL: %w", err)
	}
	_, err = url.JoinPath(sourceURL.Host, sourceURL.Path)
	if err != nil {
		return config{}, fmt.Errorf("creating source prefix from URL: %w", err)
	}
	_, err = url.Parse(conf.Host)
	if err != nil {
		return config{}, fmt.Errorf("%w failed to parse host", err)
	}

	return conf, nil
}

func (c config) log() {
	timber.Info("           host =", c.Host)
	timber.Info("       packages =", len(c.Packages))
	timber.Info("  source prefix =", c.SourcePrefix)
	if c.Favicon != "" {
		timber.Info("        favicon =", c.Favicon)
	}
	if c.RootRedirect != "" {
		timber.Info("  root redirect =", c.RootRedirect)
	}
	if c.Logs.Timezone != "" {
		timber.Info("   log timezone =", c.Logs.Timezone)
	}
	if c.Logs.TimeFormat != "" {
		timber.Info("log time format =", c.Logs.TimeFormat)
	}
}
