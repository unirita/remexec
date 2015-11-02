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

[ssh]
port             = 1234
use_certificate  = 1
private_key_file = "test.pem"
temporary_dir    = "/tmp"
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
	if c.SSH.Port != 1234 {
		t.Errorf("c.SSH.Port => %d, wants %d", c.SSH.Port, 1234)
	}
	if c.SSH.UseCertificate != 1 {
		t.Errorf("c.SSH.UseCertificate => %d, wants %d", c.SSH.UseCertificate, 1)
	}
	if c.SSH.PrivateKeyFile != "test.pem" {
		t.Errorf("c.SSH.PrivateKeyFile => %s, wants %s", c.SSH.PrivateKeyFile, "test.pem")
	}
	if c.SSH.TemporaryDir != "/tmp" {
		t.Errorf("c.SSH.TemporaryDir => %s, wants %s", c.SSH.TemporaryDir, "/tmp")
	}
}

func TestLoadReader_Normal_Empty(t *testing.T) {
	configString := ``

	c, err := loadReader(strings.NewReader(configString))
	if err != nil {
		t.Fatalf("Error occured in loadReader: %s", err)
	}
	if c.Remote.Host != "" {
		t.Errorf("c.Remote.Host => %s, must be empty", c.Remote.Host)
	}
	if c.Remote.User != "" {
		t.Errorf("c.Remote.User => %s, must be empty", c.Remote.User)
	}
	if c.Remote.Pass != "" {
		t.Errorf("c.Remote.Pass => %s, must be empty", c.Remote.Pass)
	}
	if c.Remote.IsWindows != 0 {
		t.Errorf("c.Remote.IsWindows => %d, wants %d", c.Remote.IsWindows, 0)
	}
	if c.SSH.Port != 0 {
		t.Errorf("c.SSH.Port => %d, wants %d", c.SSH.Port, 0)
	}
	if c.SSH.UseCertificate != 0 {
		t.Errorf("c.SSH.UseCertificate => %d, wants %d", c.SSH.UseCertificate, 0)
	}
	if c.SSH.PrivateKeyFile != "" {
		t.Errorf("c.SSH.PrivateKeyFile => %s, must be empty", c.SSH.PrivateKeyFile)
	}
	if c.SSH.TemporaryDir != "" {
		t.Errorf("c.SSH.TemporaryDir => %s, must be empty", c.SSH.TemporaryDir)
	}
}

func TestLoadReader_Abnormal_ParseError(t *testing.T) {
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
