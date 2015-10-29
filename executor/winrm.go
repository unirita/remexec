package executor

import "strings"

var powershellTmpArg = `"& {invoke-command -ComputerName [remoteHost]
 -Credential (ConvertTo-SecureString [pass]
 -AsPlainText -Force | % { New-Object System.Management.Automation.PSCredential([userName], $_) } | % { Get-Credential $_ })
 -ScriptBlock{Invoke-Expression $args[0]}
 -argumentList [cmd]}"`

var (
	PowershellExeAbsPath  string
	powershellExeOption   string
	pwoershellExeArgument string
)

type WinrmExecutor struct {
	host string
	user string
	pass string
}

func NewWinrmExecutor(host, user, pass string) *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = host
	e.user = user
	e.pass = pass
	return e
}

func (e *WinrmExecutor) ExecuteCommand(command string) error {
	// TODO: Call command with powershell.exe

	return nil
}

func (e *WinrmExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with powershell.exe
	return nil
}

func FetchFileAbsPath(file string) string {
	//TODO: Search file and return file abs path.

	return ""
}

func createPowershellExeArgument(host, user, pass, cmd string) string {
	//TODO: create an argument of power shell using host, user, pass.
	r := strings.NewReplacer("[remoteHost]", host, "[userName]", user, "[pass]", pass, "[cmd]", cmd)
	arg := r.Replace(powershellTmpArg)

	return arg
}
