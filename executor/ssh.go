package executor

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"

	"github.com/unirita/remexec/config"
)

const defaultPort = 22

// SSHExecutor is a Executor which executes remote command with SSH.
type SSHExecutor struct {
	config *ssh.ClientConfig
	addr   string
}

// NewSSHExecutor creates a SSHExecutor and set it client config .
func NewSSHExecutor(cfg *config.Config) *SSHExecutor {
	e := new(SSHExecutor)
	if cfg.SSH.Port <= 0 {
		e.addr = fmt.Sprintf("%s:%d", cfg.Remote.Host, defaultPort)
	} else {
		e.addr = fmt.Sprintf("%s:%d", cfg.Remote.Host, cfg.SSH.Port)
	}

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
	conn, err := ssh.Dial("tcp", e.addr, e.config)
	if err != nil {
		return fmt.Errorf("Dial error: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("Create session error: %s", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return fmt.Errorf("Request pseudo terminal error: %s", err)
	}

	command = fmt.Sprintf("%s; echo RC=$?", command)
	if err := session.Run(command); err != nil {
		return fmt.Errorf("Run command error: %s", err)
	}

	return nil
}

func (e *SSHExecutor) ExecuteScript(path string) error {
	// TODO: Execute script file with golang.org/x/crypto/ssh
	return nil
}
