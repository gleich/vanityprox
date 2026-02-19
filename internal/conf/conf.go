package conf

import "go.mattglei.ch/timber"

type Config struct {
	Host         string   `toml:"host"`
	SourcePrefix string   `toml:"source_prefix"`
	Packages     []string `toml:"packages"`
}

func (c Config) Log() {
	timber.Info("           host =", c.Host)
	timber.Info("       packages =", len(c.Packages))
	timber.Info("  source prefix =", c.SourcePrefix)
}
