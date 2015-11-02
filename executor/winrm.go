package executor

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

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

func (e *WinrmExecutor) ExecuteCommand(command string) (int, error) {
	cmd := exec.Command(powershellExeAbsPath, powershellExeFileOption, cmdExeScriptPath, e.host, e.user, e.pass, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	rc, err := e.getRC(cmdRun(cmd))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func (e *WinrmExecutor) ExecuteScript(path string) (int, error) {
	cmd := exec.Command(powershellExeAbsPath, powershellExeFileOption, localScrptExePath, e.host, e.user, e.pass, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	rc, err := e.getRC(cmdRun(cmd))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func run(cmd *exec.Cmd) error {
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (e *WinrmExecutor) getRC(err error) (int, error) {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
			return -1, err
		}
		return -1, err
	}
	return 0, nil
}
