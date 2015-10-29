// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

import "testing"

func TestNew_LocalUnixAndRemoteWindows(t *testing.T) {
	e := New("testhost", "testuser", "testpass", true)
	if _, ok := e.(*WinexeExecutor); !ok {
		t.Error("Type of executor must be WinexeExecutor, but was not.")
	}
}

func TestNew_LocalUnixAndRemoteUnix(t *testing.T) {
	e := New("testhost", "testuser", "testpass", false)
	if _, ok := e.(*SSHExecutor); !ok {
		t.Error("Type of executor must be SSHExecutor, but was not.")
	}
}
