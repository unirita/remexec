package executor

func New(host, user, pass, os string) Executor {
	switch os {
	case "windows":
		return NewWinrmExecutor(host, user, pass)
	default:
		return NewSSHExecuter(host, user, pass)
	}
}
