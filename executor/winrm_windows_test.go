package executor

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unirita/remexec/config"
)

func makeTestWinrmExecutor() *WinRMExecutor {
	e := new(WinRMExecutor)
	e.host = "host"
	e.user = "user"
	e.pass = "pass"

	return e
}

func makeWinRMExecutionSuccess() {
	cmdRun = func(*exec.Cmd) error {
		return nil
	}
}

func makeCommandFailed() {
	cmdRun = func(*exec.Cmd) error {
		return errors.New("command failed.")
	}
}
func restoreCommandFunc() {
	cmdRun = run
}

func TestNewWinrmExecutor_ValueCheck(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "host"
	c.Remote.User = "user"
	c.Remote.Pass = "pass"
	e := NewWinRMExecutor(c)

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

func TestCreateCmd_ValueCheckCmd_Normal(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	_, err := os.Stat(cmd.Args[0])
	if err != nil {
		t.Errorf("Can not access powershell.exe. [%s]", cmd.Args[0])
	}

	if cmd.Args[1] != "-File" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[1])
	}

	if !strings.Contains(cmd.Args[2], "remexec.ps1") {
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

func TestCreateCmd_ValueCheckCmd_Argument(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("dir c:\\", WINRM_CMD)

	if cmd.Args[7] != "dir c:\\" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}

	cmd = e.createCmd("c:\remote_script.ps1", WINRM_CMD)
	if cmd.Args[7] != "c:\remote_script.ps1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}

	cmd = e.createCmd("c:\remote_script.ps1 arg1", WINRM_CMD)
	if cmd.Args[7] != "c:\remote_script.ps1 arg1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}

	cmd = e.createCmd("c:\remote_script.ps1 arg1 arg2", WINRM_CMD)
	if cmd.Args[7] != "c:\remote_script.ps1 arg1 arg2" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}
}

func TestCreateCmd_ValueCheckLocalScript_Normal(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1", WINRM_LOCAL)

	_, err := os.Stat(cmd.Args[0])
	if err != nil {
		t.Errorf("Can not access powershell.exe. [%s]", cmd.Args[0])
	}

	if cmd.Args[1] != "-File" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[1])
	}

	if !strings.Contains(cmd.Args[2], "remexec.ps1") {
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

func TestCreateCmd_ValueCheckLocalScript_Argument(t *testing.T) {
	e := makeTestWinrmExecutor()
	cmd := e.createCmd("local_script.ps1 arg1", WINRM_LOCAL)

	if cmd.Args[7] != "local_script.ps1 arg1" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}

	cmd = e.createCmd("local_script.ps1 arg1 arg2", WINRM_LOCAL)
	if cmd.Args[7] != "local_script.ps1 arg1 arg2" {
		t.Errorf("The value that you expect to pass is not turned on. [%s]", cmd.Args[7])
	}

}

func TestExecuteWinRM_WinRMExecutionSuccessCmd(t *testing.T) {
	e := makeTestWinrmExecutor()
	makeWinRMExecutionSuccess()
	defer restoreCommandFunc()

	cmd := e.createCmd("ipconfig", WINRM_CMD)

	rc, err := e.executeWinRM(cmd)

	if rc != 0 {
		t.Errorf("return code => %d, wants => %d ", rc, 0)
	}

	if err != nil {
		t.Errorf("An error has occurred that is not expected. %s", err)
	}
}

func TestExecuteWinRM_WinRMExecutionSuccessLocalScript(t *testing.T) {
	e := makeTestWinrmExecutor()
	makeWinRMExecutionSuccess()
	defer restoreCommandFunc()

	cmd := e.createCmd("test.ps1", WINRM_LOCAL)

	rc, err := e.executeWinRM(cmd)

	if rc != 0 {
		t.Errorf("return code => %d, wants => %d ", rc, 0)
	}

	if err != nil {
		t.Errorf("An error has occurred that is not expected. %s", err)
	}
}

func TestExecuteWinRM_CommandFailed(t *testing.T) {
	e := makeTestWinrmExecutor()
	makeCommandFailed()
	defer restoreCommandFunc()

	cmd := e.createCmd("ipconfig", WINRM_CMD)

	rc, err := e.executeWinRM(cmd)

	if rc != -1 {
		t.Errorf("return code => %d, wants => %d ", rc, -1)
	}

	if err == nil {
		t.Errorf("An error has occurred that is not expected.")
	}

	if !strings.Contains(err.Error(), "Run command error: command failed") {
		t.Errorf("error message => %s, wants => %s ", err, "Run command error: command failed")
	}

}

//Build environment windows only
func TestExecuteWinRM_NoExistCredential(t *testing.T) {
	e := makeTestWinrmExecutor()
	e.host = "noexist"
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	cmd.Stderr = nil

	rc, err := e.executeWinRM(cmd)

	if err != nil {
		t.Fatalf("An error has occurred that is not expected. %s", err)
	}

	if rc != 250 {
		t.Errorf("return code => %d, wants => %d ", rc, 250)
	}
}
