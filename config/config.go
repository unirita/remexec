// Package config provides access interface for config file.
package config

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

// Config has all sections and items in config file.
type Config struct {
	Remote RemoteSection
}

// RemoteSection is [remote] section in config file.
type RemoteSection struct {
	Host      string `toml:"host"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
	IsWindows int    `toml:"is_windows"`
}

// Load loads config from file which is in path.
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return loadReader(f)
}

// loadReader loads config from reader.
func loadReader(reader io.Reader) (*Config, error) {
	c := new(Config)
	if _, err := toml.DecodeReader(reader, c); err != nil {
		return nil, err
	}
	return c, nil
}
