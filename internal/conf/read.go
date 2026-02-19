package conf

import (
	"fmt"
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

	return c, nil
}
