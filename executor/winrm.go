package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/unirita/remexec/config"
)

const option = "-File"

type WinrmExecutor struct {
	host       string
	user       string
	pass       string
	powershell string
	winrm      string
}

const (
	WINRM_CMD   = "e"
	WINRM_LOCAL = "f"
)

type commandRunFunc func(*exec.Cmd) error

var cmdRun commandRunFunc = run

func NewWinrmExecutor(cfg *config.Config) *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	e.winrm = cfg.WinRM.WinRMScriptPath

	return e
}

func (e *WinrmExecutor) ExecuteCommand(command string) (int, error) {
	return e.ExecuteWinRM(command, WINRM_CMD)
}

func (e *WinrmExecutor) ExecuteScript(path string) (int, error) {
	return e.ExecuteWinRM(path, WINRM_LOCAL)
}

func (e *WinrmExecutor) ExecuteWinRM(execution string, scripttype string) (int, error) {
	cmd := e.createCmd(execution, scripttype)

	rc, err := e.getRC(cmdRun(cmd))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func (e *WinrmExecutor) createCmd(execution string, scripttype string) *exec.Cmd {
	cmd := new(exec.Cmd)

	powershell := strings.Replace(os.Getenv("PSModulePath"), "Modules\\", "powershell.exe", -1)

	cmd = exec.Command(powershell, option, e.winrm, scripttype, e.host, e.user, e.pass, execution)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd

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
