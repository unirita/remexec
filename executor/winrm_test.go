package executor

import (
	"errors"
	"os/exec"
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

func makeCmmandSuccess() {
	cmdRun = func(*exec.Cmd) error {
		return nil
	}
}

func restCommandFunc() {
	cmdRun = run
}

func makeCommandFailed() {
	cmdRun = func(*exec.Cmd) error {
		return errors.New("error")
	}
}

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

func TestExecuteCommand_SuccessCommand(t *testing.T) {
	e := makeTestWinrmExecutor()

	makeCmmandSuccess()
	defer restCommandFunc()

	if err := e.ExecuteCommand(""); err != nil {
		t.Errorf("Error occurs that is not expected.")
	}
}

func TestExecuteCommand_FailedCommand(t *testing.T) {
	e := makeTestWinrmExecutor()

	makeCommandFailed()
	defer restCommandFunc()

	if err := e.ExecuteCommand(""); err == nil {
		t.Errorf("Error did not occur.")
	}
}

func TestExecuteScript_ExecutePowershellScript(t *testing.T) {

}

func TestCreatecreatePSCommandArgument_ValueCheck(t *testing.T) {
	expect := "& {invoke-command -ComputerName \"hostName\" -Credential (ConvertTo-SecureString \"password\" -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential(\"userName\", $_) } | % { Get-Credential $_ }) -ScriptBlock{Invoke-Expression $args[0]} -argumentList \"echo hoge \"}; echo $?"

	result := createPSCommandArgument("hostName", "userName", "password", "echo hoge")

	if result != expect {
		t.Errorf("It different from the contents of result is expecting. [%s]", result)
	}

}

func TestCreatePSScriptArgument_ValueCheck(t *testing.T) {
	expect := "& {invoke-command -ComputerName \"hostName\" -Credential (ConvertTo-SecureString \"password\" -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential(\"userName\", $_) } | % { Get-Credential $_ }) -File \"script.ps1\" }"

	result := createPSScriptArgument("hostName", "userName", "password", "script.ps1")

	if result != expect {
		t.Errorf("It different from the contents of result is expecting. [%s]", result)
	}

}
