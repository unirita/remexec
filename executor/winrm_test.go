package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

func makeTestWinrmExecutor() *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = "host"
	e.user = "user"
	e.pass = "pass"

	return e
}

//func makeCmmandSuccess() {
//	cmdRun = func(*exec.Cmd) error {
//		return nil
//	}
//}

//func restCommandFunc() {
//	cmdRun = run
//}

//func makeCommandFailed() {
//	cmdRun = func(*exec.Cmd) error {
//		return errors.New("error")
//	}
//}

func TestNewWinrmExecutor_ValueCheck(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "host"
	c.Remote.User = "user"
	c.Remote.Pass = "pass"
	e := NewWinrmExecutor(c)

	if e.host != "host" {
		t.Errorf("The value that you expect to host is not turned on. [%s]", e.host)
	}

	if e.user != "user" {
		t.Errorf("The value that you expect to user is not turned on. [%s]", e.user)
	}

	if e.pass != "pass" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", e.pass)
	}

}
