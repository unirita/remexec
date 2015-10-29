package executor

import (
	"os/exec"
	"strings"

	"github.com/unirita/remexec/config"
)

var psExeTmpArgByCmd = `"& {invoke-command -ComputerName [remoteHost]
 -Credential (ConvertTo-SecureString [pass]
 -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential([userName], $_) } | % { Get-Credential $_ })
 -ScriptBlock{Invoke-Expression $args[0]}
 -argumentList [cmd]}"`

var psExeTmpArgByScript = `"& {invoke-command -ComputerName [remoteHost]
 -Credential (ConvertTo-SecureString [pass]
 -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential([userName], $_) } | % { Get-Credential $_ })
 -File [script] }"`

const powershellExeAbsPath = "powershell.exe"
const powershellExeOption = "-Command"

type WinrmExecutor struct {
	host string
	user string
	pass string
}

func NewWinrmExecutor(cfg *config.Config) *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	return e
}

func (e *WinrmExecutor) ExecuteCommand(command string) error {
	// TODO: Call command with powershell.exe
	cmdArg := createPSCommandArgument(e.host, e.user, e.pass, command)
	cmd := exec.Command(powershellExeAbsPath, powershellExeOption, cmdArg)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (e *WinrmExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with powershell.exe
	return nil
}

func createPSCommandArgument(host, user, pass, cmd string) string {
	r := strings.NewReplacer("[remoteHost]", host, "[userName]", user, "[pass]", pass, "[cmd]", cmd)
	arg := r.Replace(psExeTmpArgByCmd)

	return arg
}

func createPSScriptArgument(host, user, pass, script string) string {
	r := strings.NewReplacer("[remoteHost]", host, "[userName]", user, "[pass]", pass, "[script]", script)
	arg := r.Replace(psExeTmpArgByScript)

	return arg
}
