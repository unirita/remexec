package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

func TestNewSSHExecutor_Password(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.SSH.UseCertificate = 0

	e := NewSSHExecutor(c)
	if e.host != "testhost" {
		t.Errorf("executor.host => %s, wants %s", e.host, "testhost")
	}
	if e.config.User != "testuser" {
		t.Errorf("executor.config.User => %s, wants %s", e.config.User, "testuser")
	}
	if len(e.config.Auth) != 1 {
		t.Fatalf("len(executor.config.Auth) => %d, wants %d", len(e.config.Auth), 1)
	}
	if e.config.Auth[0] == nil {
		t.Error("executor.config.Auth[0] must be set, but it was not.")
	}
}

func TestNewSSHExecutor_Certificate_PrivateKeyNotFound(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.SSH.UseCertificate = 1
	c.SSH.PrivateKeyFile = "noexists"

	e := NewSSHExecutor(c)
	if e.host != "testhost" {
		t.Errorf("executor.host => %s, wants %s", e.host, "testhost")
	}
	if e.config.User != "testuser" {
		t.Errorf("executor.config.User => %s, wants %s", e.config.User, "testuser")
	}
	if len(e.config.Auth) != 1 {
		t.Fatalf("len(executor.config.Auth) => %d, wants %d", len(e.config.Auth), 1)
	}
	if e.config.Auth[0] != nil {
		t.Error("executor.config.Auth[0] must not be set, but it was.")
	}
}
