package executor

import (
	"os"
	"testing"

	"github.com/unirita/remexec/config"
)

func makeTestWinrmExecutor() *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = "host"
	e.user = "user"
	e.pass = "pass"
	e.winrm = "winrm.ps1"

	return e
}

//func makeCmmandSuccess() {
//	cmdRun = func(*exec.Cmd) error {
//		return nil
//	}
//}

//func restoreCommandFunc() {
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

	if e.winrm != "winrm.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", e.winrm)
	}

}

func TestCreateCmd_ValueCheckCmd_ArgumentOne(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"

	if cmd.Args[0] != path {
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

func TestCreateCmd_ValueCheckCmd_ArgumentTwo(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("dir c:\\", WINRM_CMD)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"

	if cmd.Args[0] != path {
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

	if cmd.Args[7] != "dir c:\\" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckCmd_ArgumentThree(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("c:\\test.ps1 arg1 arg2", WINRM_CMD)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"

	if cmd.Args[0] != path {
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

	if cmd.Args[7] != "c:\\test.ps1 arg1 arg2" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckLocalScript_NoArgument(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1", WINRM_LOCAL)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"
	if cmd.Args[0] != path {
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

func TestCreateCmd_ValueCheckLocalScript_ArgumentOne(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1 arg1", WINRM_LOCAL)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"
	if cmd.Args[0] != path {
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

	if cmd.Args[7] != "local_script.ps1 arg1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckLocalScript_ArgumentTwo(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1 arg1 arg2", WINRM_LOCAL)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"
	if cmd.Args[0] != path {
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

	if cmd.Args[7] != "local_script.ps1 arg1 arg2" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckLocalScript_ArgumentThree(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1 arg1 arg2 arg3", WINRM_LOCAL)

	path := os.Getenv("PSModulePath")
	path += "powershell.exe"
	if cmd.Args[0] != path {
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

	if cmd.Args[7] != "local_script.ps1 arg1 arg2 arg3" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}
