package executor

func New(host, user, pass string, isWindows bool) Executor {
	if isWindows {
		return NewWinrmExecutor(host, user, pass)
	}
	return NewSSHExecutor(host, user, pass)
}
