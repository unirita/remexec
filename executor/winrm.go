package executor

var (
	powershellExeAbsPath  string
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
