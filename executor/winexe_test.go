package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

func TestNewWinexeExecutor(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	e := NewWinexeExecutor(c)

	if e.host != "testhost" {
		t.Errorf("e.host => %s, wants %s", e.host, "testhost")
	}
	if e.user != "testuser" {
		t.Errorf("e.user => %s, wants %s", e.user, "testuser")
	}
	if e.pass != "testpass" {
		t.Errorf("e.pass => %s, wants %s", e.pass, "testpass")
	}
}
