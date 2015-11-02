package executor

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/unirita/remexec/config"
)

type WinexeExecutor struct {
	host string
	user string
	pass string
}

func NewWinexeExecutor(cfg *config.Config) *WinexeExecutor {
	e := new(WinexeExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	return e
}

func (e *WinexeExecutor) ExecuteCommand(command string) error {
	userParam := fmt.Sprintf("%s%%%s", e.user, e.pass)
	hostParam := fmt.Sprintf("//%s", e.host)
	commandParam := fmt.Sprintf("cmd /c %s", command)
	cmd := exec.Command("winexe", "-U", userParam, hostParam, commandParam)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	rc, err := e.getRC(cmd.Run())
	if err != nil {
		return fmt.Errorf("Run command error: %s", err)
	}
	fmt.Println("RC = %d", rc)

	return nil
}

func (e *WinexeExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with winexe
	return nil
}

func (e *WinexeExecutor) getRC(err error) (int, error) {
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				return s.ExitStatus(), nil
			}
		}
		return -1, err
	}
	return 0, nil
}
