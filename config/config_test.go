package config

import (
	"strings"
	"testing"
)

func TestLoad_FileNotExists(t *testing.T) {
	_, err := Load("noexists")
	if err == nil {
		t.Error("Error must be occured, but was not.")
	}
}

func TestLoadReader_Normal(t *testing.T) {
	configString := `
[remote]
host       = "testhost"
user       = "testuser"
pass       = "testpass"
is_windows = 1
`

	c, err := loadReader(strings.NewReader(configString))
	if err != nil {
		t.Fatalf("Error occured in loadReader: %s", err)
	}
	if c.Remote.Host != "testhost" {
		t.Errorf("c.Remote.Host => %s, wants %s", c.Remote.Host, "testhost")
	}
	if c.Remote.User != "testuser" {
		t.Errorf("c.Remote.User => %s, wants %s", c.Remote.User, "testuser")
	}
	if c.Remote.Pass != "testpass" {
		t.Errorf("c.Remote.Pass => %s, wants %s", c.Remote.Pass, "testpass")
	}
	if c.Remote.IsWindows != 1 {
		t.Errorf("c.Remote.IsWindows => %d, wants %d", c.Remote.IsWindows, 1)
	}
}

func TestLoadReader_ParseError(t *testing.T) {
	configString := `
[remote
host       = "testhost"
user       = "testuser"
pass       = "testpass"
is_windows = 1
`

	_, err := loadReader(strings.NewReader(configString))
	if err == nil {
		t.Error("Error must be occured, but was not.")
	}
}
