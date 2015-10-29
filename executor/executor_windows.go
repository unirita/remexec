package executor

import "github.com/unirita/remexec/config"

func New(cfg *config.Config) Executor {
	if cfg.Remote.IsWindows != 0 {
		return NewWinrmExecutor(cfg)
	}
	return NewSSHExecutor(cfg)
}
