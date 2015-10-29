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

func TestLoadReader_Normal_Windows(t *testing.T) {
	c := `
[remote]
host       = "testhost"
user       = "testuser"
pass       = "testpass"
is_windows = 1
`

	err := loadReader(strings.NewReader(c))
	if err != nil {
		t.Fatalf("Error occured in loadReader: %s", err)
	}
	if Host != "testhost" {
		t.Errorf("Host => %s, wants %s", Host, "testhost")
	}
	if User != "testuser" {
		t.Errorf("User => %s, wants %s", User, "testuser")
	}
	if Pass != "testpass" {
		t.Errorf("Pass => %s, wants %s", Pass, "testpass")
	}
	if !IsWindows {
		t.Errorf("IsWindows must be true, but it was not.")
	}
}

func TestLoadReader_Normal_NotWindows(t *testing.T) {
	c := `
[remote]
host       = "testhost"
user       = "testuser"
pass       = "testpass"
is_windows = 0
`

	err := loadReader(strings.NewReader(c))
	if err != nil {
		t.Fatalf("Error occured in loadReader: %s", err)
	}
	if IsWindows {
		t.Errorf("IsWindows must be false, but it was not.")
	}
}

func TestLoadReader_ParseError(t *testing.T) {
	c := `
[remote
host       = "testhost"
user       = "testuser"
pass       = "testpass"
is_windows = 1
`

	err := loadReader(strings.NewReader(c))
	if err == nil {
		t.Error("Error must be occured, but was not.")
	}
}
