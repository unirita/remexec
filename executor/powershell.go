package executor

type PowershellExecuter struct {
	host string
	user string
	pass string
}

func NewPowershellExecuter(host, user, pass string) *PowershellExecuter {
	e := new(PowershellExecuter)
	e.host = host
	e.user = user
	e.pass = pass
	return e
}

func (e *PowershellExecuter) ExecuteCommand(command string) error {
	// TODO: Call command with powershell.exe
	return nil
}

func (e *PowershellExecuter) ExecuteScript(path string) error {
	// TODO: Execute script file with powershell.exe
	return nil
}
