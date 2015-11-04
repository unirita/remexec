package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

func TestNew_LocalWindowsAndRemoteWindows(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.Remote.IsWindows = 1
	e := New(c)
	if _, ok := e.(*WinRMExecutor); !ok {
		t.Error("Type of executor must be PowershellExecuter, but was not.")
	}
}

func TestNew_LocalWindowsAndRemoteUnix(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.Remote.IsWindows = 0
	e := New(c)
	if _, ok := e.(*SSHExecutor); !ok {
		t.Error("Type of executor must be SSHExecuter, but was not.")
	}
}
