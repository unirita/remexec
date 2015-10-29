// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

func TestNew_LocalUnixAndRemoteWindows(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.Remote.IsWindows = 1
	e := New(c)
	if _, ok := e.(*WinexeExecutor); !ok {
		t.Error("Type of executor must be WinexeExecutor, but was not.")
	}
}

func TestNew_LocalUnixAndRemoteUnix(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.Remote.IsWindows = 0
	e := New(c)
	if _, ok := e.(*SSHExecutor); !ok {
		t.Error("Type of executor must be SSHExecutor, but was not.")
	}
}
