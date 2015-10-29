package executor

import "testing"

func TestNew_LocalWindowsAndRemoteWindows(t *testing.T) {
	e := New("testhost", "testuser", "testpass", true)
	if _, ok := e.(*WinrmExecutor); !ok {
		t.Error("Type of executor must be PowershellExecuter, but was not.")
	}
}

func TestNew_LocalWindowsAndRemoteUnix(t *testing.T) {
	e := New("testhost", "testuser", "testpass", false)
	if _, ok := e.(*SSHExecuter); !ok {
		t.Error("Type of executor must be SSHExecuter, but was not.")
	}
}
