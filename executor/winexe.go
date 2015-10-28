package executor

type WinexeExecuter struct {
	host string
	user string
	pass string
}

func NewWinexeExecuter(host, user, pass string) *WinexeExecuter {
	e := new(WinexeExecuter)
	e.host = host
	e.user = user
	e.pass = pass
	return e
}

func (e *WinexeExecuter) ExecuteCommand(command string) error {
	// TODO: Call command with winexe
	return nil
}

func (e *WinexeExecuter) ExecuteScript(path string) error {
	// TODO: Execute script file with winexe
	return nil
}
