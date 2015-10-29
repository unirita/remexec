package executor

import "github.com/unirita/remexec/config"

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
	// TODO: Call command with winexe
	return nil
}

func (e *WinexeExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with winexe
	return nil
}
