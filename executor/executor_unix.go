// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

func New(host, user, pass, os string) Executor {
	switch os {
	case "windows":
		return NewWinexeExecuter(host, user, pass)
	default:
		return NewSSHExecuter(host, user, pass)
	}
}
