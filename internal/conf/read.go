package conf

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Read() (Config, error) {
	filename := "config.toml"
	bin, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("reading from %s: %w", filename, err)
	}

	var c Config
	err = toml.Unmarshal(bin, &c)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshaling toml: %w", err)
	}

	// ensure that source prefix is formatted properly
	sourceURL, err := url.Parse(c.SourcePrefix)
	if err != nil {
		return Config{}, fmt.Errorf("parsing source prefix URL: %w", err)
	}
	c.SourcePrefix, err = url.JoinPath(sourceURL.Host, sourceURL.Path)
	if err != nil {
		return Config{}, fmt.Errorf("creating source prefix from URL: %w", err)
	}
	hostURL, err := url.Parse(c.Host)
	if err != nil {
		return Config{}, fmt.Errorf("%w failed to parse host", err)
	}
	c.Host = hostURL.Host

	return c, nil
}
