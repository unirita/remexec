package executor

import (
	"testing"

	"github.com/unirita/remexec/config"
)

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

func TestExecuteCommand_ExecuteCommandPowershell(t *testing.T) {

}

func TestExecuteScript_ExecutePowershellScript(t *testing.T) {

}

func TestCreatecreatePSCommandArgument_ValueCheck(t *testing.T) {
	expect := "& {invoke-command -ComputerName \"hostName\" -Credential (ConvertTo-SecureString \"password\" -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential(\"userName\", $_) } | % { Get-Credential $_ }) -ScriptBlock{Invoke-Expression $args[0]} -argumentList \"echo hoge \"}"

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
