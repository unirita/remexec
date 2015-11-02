package executor

import (
	"os"
	"os/exec"

	"github.com/unirita/remexec/config"
)

const (
	powershellExeAbsPath    = "C:\\WINDOWS\\system32\\WindowsPowerShell\\v1.0\\powershell.exe"
	powershellExeFileOption = "-File"
	cmdExeScriptPath        = "script\\remotecmdexe.ps1"
	localScrptExePath       = "script\\localscriptexe.ps1"
)

type WinrmExecutor struct {
	host string
	user string
	pass string
}

type commandRunFunc func(*exec.Cmd) error

var cmdRun commandRunFunc = run

func NewWinrmExecutor(cfg *config.Config) *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	return e
}

func (e *WinrmExecutor) ExecuteCommand(command string) error {
	cmd := exec.Command(powershellExeAbsPath, powershellExeFileOption, cmdExeScriptPath, e.host, e.user, e.pass, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmdRun(cmd)
	if err != nil {
		return err
	}

	return nil
}

func run(cmd *exec.Cmd) error {
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (e *WinrmExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with powershell.exe
	cmd := exec.Command(powershellExeAbsPath, powershellExeFileOption, localScrptExePath, e.host, e.user, e.pass, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmdRun(cmd)
	if err != nil {
		return err
	}

	return nil
}
