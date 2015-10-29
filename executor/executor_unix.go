// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

import "github.com/unirita/remexec/config"

func New(cfg *config.Config) Executor {
	if cfg.Remote.IsWindows != 0 {
		return NewWinexeExecutor(cfg)
	}
	return NewSSHExecutor(cfg)
}
