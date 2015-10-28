// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

import "testing"

func TestNew_LocalUnixAndRemoteWindows(t *testing.T) {
	e := New("testhost", "testuser", "testpass", "windows")
	if _, ok := e.(*WinexeExecuter); !ok {
		t.Error("Type of executor must be WinexeExecutor, but was not.")
	}
}

func TestNew_LocalUnixAndRemoteUnix(t *testing.T) {
	e := New("testhost", "testuser", "testpass", "linux")
	if _, ok := e.(*SSHExecuter); !ok {
		t.Error("Type of executor must be SSHExecuter, but was not.")
	}
}
