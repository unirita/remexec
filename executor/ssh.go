package executor

import "github.com/unirita/remexec/config"

type SSHExecutor struct {
	host string
	user string
	pass string
}

func NewSSHExecutor(cfg *config.Config) *SSHExecutor {
	e := new(SSHExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass
	return e
}

func (e *SSHExecutor) ExecuteCommand(command string) error {
	// TODO: Call command with golang.org/x/crypto/ssh
	return nil
}

func (e *SSHExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with golang.org/x/crypto/ssh
	return nil
}
