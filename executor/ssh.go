package executor

type SSHExecutor struct {
	host string
	user string
	pass string
}

func NewSSHExecutor(host, user, pass string) *SSHExecutor {
	e := new(SSHExecutor)
	e.host = host
	e.user = user
	e.pass = pass
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
