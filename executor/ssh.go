package executor

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/unirita/remexec/config"
)

const (
	defaultPort   = 22
	defaultTmpDir = "/tmp"
)

// SSHExecutor is a Executor which executes remote command with SSH.
type SSHExecutor struct {
	config *ssh.ClientConfig
	addr   string
	tmpDir string
}

// NewSSHExecutor creates a SSHExecutor and set it client config .
func NewSSHExecutor(cfg *config.Config) *SSHExecutor {
	e := new(SSHExecutor)
	if cfg.SSH.Port <= 0 {
		e.addr = fmt.Sprintf("%s:%d", cfg.Remote.Host, defaultPort)
	} else {
		e.addr = fmt.Sprintf("%s:%d", cfg.Remote.Host, cfg.SSH.Port)
	}
	e.tmpDir = cfg.SSH.TemporaryDir
	if e.tmpDir == "" {
		e.tmpDir = defaultTmpDir
	}
	fmt.Println(cfg.SSH.Port)
	fmt.Println(cfg.SSH.UseCertificate)
	fmt.Println(cfg.SSH.PrivateKeyFile)
	fmt.Println(cfg.SSH.TemporaryDir)

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

func (e *SSHExecutor) ExecuteCommand(command string) (int, error) {
	conn, err := ssh.Dial("tcp", e.addr, e.config)
	if err != nil {
		return -1, fmt.Errorf("Dial error: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return -1, fmt.Errorf("Create session error: %s", err)
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
		return -1, fmt.Errorf("Request pseudo terminal error: %s", err)
	}

	rc, err := e.getRC(session.Run(command))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func (e *SSHExecutor) ExecuteScript(pathWithParam string) (int, error) {
	path, param := splitPathAndParam(pathWithParam)
	remotePath := generateRemotePath(path, e.tmpDir)
	command, err := generateCreateCommand(path, remotePath)
	if err != nil {
		return -1, err
	}
	if _, err := e.ExecuteCommand(command); err != nil {
		return -1, err
	}
	defer e.ExecuteCommand(generateCleanCommand(remotePath))

	return e.ExecuteCommand(generateExecuteCommand(remotePath, param))
}

func (e *SSHExecutor) getRC(err error) (int, error) {
	if err != nil {
		if e2, ok := err.(*ssh.ExitError); ok {
			return e2.ExitStatus(), nil
		}
		return -1, err
	}
	return 0, nil
}

func splitPathAndParam(pathWithParam string) (string, string) {
	pathBuf := strings.SplitN(pathWithParam, " ", 2)
	path := pathBuf[0]
	param := ""
	if len(pathBuf) > 1 {
		param = pathBuf[1]
	}

	return path, param
}

func generateRemotePath(localPath, tmpDir string) string {
	name := fmt.Sprintf("%s/%s.%s", tmpDir,
		time.Now().Format("20060102150405.000"), filepath.Base(localPath))
	return name
}

func generateCreateCommand(localPath, remotePath string) (string, error) {
	script, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("Open script file error: %s", err)
	}
	defer script.Close()

	commandBuf := new(bytes.Buffer)
	s := bufio.NewScanner(script)

	commandBuf.WriteString(fmt.Sprintf("tee %s > /dev/null << EOF\n", remotePath))
	for s.Scan() {
		commandBuf.WriteString(strings.Replace(s.Text(), `$`, `\$`, -1))
		commandBuf.WriteByte('\n')
	}
	commandBuf.WriteString("EOF\n")
	return commandBuf.String(), nil
}

func generateExecuteCommand(remotePath, param string) string {
	return fmt.Sprintf("chmod +x %s; %s %s", remotePath, remotePath, param)
}

func generateCleanCommand(remotePath string) string {
	return fmt.Sprintf("rm -f %s", remotePath)
}
