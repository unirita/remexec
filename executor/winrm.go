package executor

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/unirita/remexec/config"
)

const option = "-File"

type WinrmExecutor struct {
	host       string
	user       string
	pass       string
	powershell string
	winrmCmd   string
	winrmLocal string
}

type commandRunFunc func(*exec.Cmd) error

var cmdRun commandRunFunc = run

func NewWinrmExecutor(cfg *config.Config) *WinrmExecutor {
	e := new(WinrmExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	e.powershell = cfg.WinRM.PowershellPath
	e.winrmCmd = cfg.WinRM.WinRMCmdPath
	e.winrmLocal = cfg.WinRM.WinRMLocalPath
	return e
}

func (e *WinrmExecutor) ExecuteCommand(command string) (int, error) {
	cmd := exec.Command(e.powershell, option, e.winrmCmd, e.host, e.user, e.pass, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	rc, err := e.getRC(cmdRun(cmd))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func (e *WinrmExecutor) ExecuteScript(path string) (int, error) {
	cmd := exec.Command(e.powershell, option, e.winrmLocal, e.host, e.user, e.pass, path)
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
