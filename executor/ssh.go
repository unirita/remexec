package executor

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"

	"github.com/unirita/remexec/config"
)

// SSHExecutor is a Executor which executes remote command with SSH.
type SSHExecutor struct {
	config *ssh.ClientConfig
	host   string
}

// NewSSHExecutor creates a SSHExecutor and set it client config .
func NewSSHExecutor(cfg *config.Config) *SSHExecutor {
	e := new(SSHExecutor)
	e.host = cfg.Remote.Host
	e.config = new(ssh.ClientConfig)
	e.config.User = cfg.Remote.User
	if cfg.SSH.UseCertificate != 0 {
		e.config.Auth = []ssh.AuthMethod{
			publicKeyFile(cfg.SSH.PrivateKeyFile),
		}
	} else {
		e.config.Auth = []ssh.AuthMethod{
			ssh.Password(cfg.Remote.Pass),
		}
	}

	return e
}

func publicKeyFile(file string) ssh.AuthMethod {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func (e *SSHExecutor) ExecuteCommand(command string) error {
	// TODO: Call command with golang.org/x/crypto/ssh
	return nil
}

func (e *SSHExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with golang.org/x/crypto/ssh
	return nil
}
