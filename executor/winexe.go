package executor

type WinexeExecutor struct {
	host string
	user string
	pass string
}

func NewWinexeExecutor(host, user, pass string) *WinexeExecutor {
	e := new(WinexeExecutor)
	e.host = host
	e.user = user
	e.pass = pass
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
