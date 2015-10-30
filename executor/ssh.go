package executor

import (
	"bufio"
	"bytes"
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

	rc, err := getRC(session.Run(command))
	if err != nil {
		return fmt.Errorf("Run command error: %s", err)
	}
	fmt.Printf("RC = %d\n", rc)

	return nil
}

func (e *SSHExecutor) ExecuteScript(path string) error {
	command, err := scriptToCommand(path)
	if err != nil {
		return err
	}

	return e.ExecuteCommand(command)
}

func scriptToCommand(path string) (string, error) {
	script, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Open script file error: %s", err)
	}
	defer script.Close()

	commandBuf := new(bytes.Buffer)
	s := bufio.NewScanner(script)

	commandBuf.WriteString("bash -s << EOF\n")
	for s.Scan() {
		commandBuf.Write(s.Bytes())
		commandBuf.WriteByte('\n')
	}
	commandBuf.WriteString("EOF\n")
	return commandBuf.String(), nil
}

func getRC(err error) (int, error) {
	if err != nil {
		if e2, ok := err.(*ssh.ExitError); ok {
			return e2.ExitStatus(), nil
		}
		return -1, err
	}
	return 0, nil
}
