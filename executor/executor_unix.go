// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

func New(host, user, pass string, isWindows bool) Executor {
	if isWindows {
		return NewWinexeExecuter(host, user, pass)
	}
	return NewSSHExecuter(host, user, pass)
}
