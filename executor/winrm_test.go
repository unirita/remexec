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
	e.powershell = "powershell.exe"
	e.winrm = "winrm.ps1"

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
	c.WinRM.PowershellPath = "powershell.exe"
	c.WinRM.WinRMScriptPath = "winrm.ps1"
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

	if e.powershell != "powershell.exe" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", e.powershell)
	}

	if e.winrm != "winrm.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", e.winrm)
	}

}

func TestCreateCmd_ValueCheckLocalScriptPattern(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	if cmd.Args[0] != "powershell.exe" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[0])
	}

	if cmd.Args[1] != option {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[1])
	}

	if cmd.Args[2] != "winrm.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[2])
	}

	if cmd.Args[3] != WINRM_CMD {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[3])
	}

	if cmd.Args[4] != "host" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[4])
	}

	if cmd.Args[5] != "user" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[5])
	}

	if cmd.Args[6] != "pass" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[6])
	}

	if cmd.Args[7] != "ipconfig" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckCmdPattern(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1", WINRM_LOCAL)

	if cmd.Args[0] != "powershell.exe" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[0])
	}

	if cmd.Args[1] != option {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[1])
	}

	if cmd.Args[2] != "winrm.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[2])
	}

	if cmd.Args[3] != WINRM_LOCAL {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[3])
	}

	if cmd.Args[4] != "host" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[4])
	}

	if cmd.Args[5] != "user" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[5])
	}

	if cmd.Args[6] != "pass" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[6])
	}

	if cmd.Args[7] != "local_script.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}
