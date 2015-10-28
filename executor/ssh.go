package executor

type SSHExecuter struct {
	host string
	user string
	pass string
}

func NewSSHExecuter(host, user, pass string) *SSHExecuter {
	e := new(SSHExecuter)
	e.host = host
	e.user = user
	e.pass = pass
	return e
}

func (e *SSHExecuter) ExecuteCommand(command string) error {
	// TODO: Call command with golang.org/x/crypto/ssh
	return nil
}

func (e *SSHExecuter) ExecuteScript(path string) error {
	// TODO: Execute script file with golang.org/x/crypto/ssh
	return nil
}
