package config

import (
	"strings"
	"testing"
)

func TestLoad_FileNotExists(t *testing.T) {
	err := Load("noexists")
	if err == nil {
		t.Error("Error must be occured, but was not.")
	}
}

func TestLoadReader_Normal(t *testing.T) {
	c := `
[remote]
host = "testhost"
user = "testuser"
pass = "testpass"
os   = "testos"
`

	err := loadReader(strings.NewReader(c))
	if err != nil {
		t.Fatalf("Error occured in loadReader: %s", err)
	}
	if Remote.Host != "testhost" {
		t.Errorf("Remote.Host => %s, wants %s", Remote.Host, "testhost")
	}
	if Remote.User != "testuser" {
		t.Errorf("Remote.User => %s, wants %s", Remote.User, "testuser")
	}
	if Remote.Pass != "testpass" {
		t.Errorf("Remote.Pass => %s, wants %s", Remote.Pass, "testpass")
	}
	if Remote.OS != "testos" {
		t.Errorf("Remote.OS => %s, wants %s", Remote.OS, "testos")
	}
}

func TestLoadReader_ParseError(t *testing.T) {
	c := `
[remote
host = "testhost"
user = "testuser"
pass = "testpass"
os   = "testos"
`

	err := loadReader(strings.NewReader(c))
	if err == nil {
		t.Error("Error must be occured, but was not.")
	}
}
