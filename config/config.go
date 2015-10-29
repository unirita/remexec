// Package config provides access interface for config file.
package config

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Remote remoteSection
}

// remoteSection is [remote] section in config file.
type remoteSection struct {
	Host      string `toml:"host"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
	IsWindows int    `toml:"is_windows"`
}

var (
	Host      string
	User      string
	Pass      string
	IsWindows bool
)

// Load loads config from file which is in path.
func Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return loadReader(f)
}

// loadReader loads config from reader.
func loadReader(reader io.Reader) error {
	c := new(config)
	if _, err := toml.DecodeReader(reader, c); err != nil {
		return err
	}

	Host = c.Remote.Host
	User = c.Remote.User
	Pass = c.Remote.Pass
	IsWindows = c.Remote.IsWindows != 0
	return nil
}
