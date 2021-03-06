package executor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/unirita/remexec/config"
)

func TestNewSSHExecutor_Password(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.SSH.Port = 1234
	c.SSH.UseCertificate = 0

	e := NewSSHExecutor(c)
	if e.addr != "testhost:1234" {
		t.Errorf("executor.addr => %s, wants %s", e.addr, "testhost:1234")
	}
	if e.config.User != "testuser" {
		t.Errorf("executor.config.User => %s, wants %s", e.config.User, "testuser")
	}
	if len(e.config.Auth) != 1 {
		t.Fatalf("len(executor.config.Auth) => %d, wants %d", len(e.config.Auth), 1)
	}
	if e.config.Auth[0] == nil {
		t.Error("executor.config.Auth[0] must be set, but it was not.")
	}
}

func TestNewSSHExecutor_Certificate_PrivateKeyNotFound(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.SSH.Port = 1234
	c.SSH.UseCertificate = 1
	c.SSH.PrivateKeyFile = "noexists"

	e := NewSSHExecutor(c)
	if e.addr != "testhost:1234" {
		t.Errorf("executor.addr => %s, wants %s", e.addr, "testhost:1234")
	}
	if e.config.User != "testuser" {
		t.Errorf("executor.config.User => %s, wants %s", e.config.User, "testuser")
	}
	if len(e.config.Auth) != 1 {
		t.Fatalf("len(executor.config.Auth) => %d, wants %d", len(e.config.Auth), 1)
	}
	if e.config.Auth[0] != nil {
		t.Error("executor.config.Auth[0] must not be set, but it was.")
	}
}

func TestNewSSHExecutor_DefaltPort(t *testing.T) {
	c := new(config.Config)
	c.Remote.Host = "testhost"
	c.Remote.User = "testuser"
	c.Remote.Pass = "testpass"
	c.SSH.UseCertificate = 0

	e := NewSSHExecutor(c)
	if e.addr != "testhost:22" {
		t.Errorf("executor.addr => %s, wants %s", e.addr, "testhost:22")
	}
}

func TestSplitPathWithParam_Normal(t *testing.T) {
	path, param := splitPathAndParam("/home/gocuto/test.sh test1 test2 test3")
	if path != "/home/gocuto/test.sh" {
		t.Errorf("path => %s, wants %s", path, "/home/gocuto/test.sh")
	}
	if param != "test1 test2 test3" {
		t.Errorf("param => %s, wants %s", param, "test1 test2 test3")
	}
}

func TestSplitPathWithParam_NoParam(t *testing.T) {
	path, param := splitPathAndParam("/home/gocuto/test.sh")
	if path != "/home/gocuto/test.sh" {
		t.Errorf("path => %s, wants %s", path, "/home/gocuto/test.sh")
	}
	if param != "" {
		t.Errorf("param => %s, must be empty but was not.")
	}
}

func TestGenerateCreateCommand_Normal(t *testing.T) {
	expected := `tee remote.sh > /dev/null << EOF
#!/bin/sh

echo "test message."
echo \$1
exit 1
EOF
`

	scriptPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita",
		"remexec", "executor", "_testdata", "sshtest.sh")
	result, err := generateCreateCommand(scriptPath, "remote.sh")
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if result != expected {
		t.Errorf("Convert result is not expected value.")
		t.Log("Expected:")
		t.Log(expected)
		t.Log("Actual:")
		t.Log(result)
	}
}

func TestGenerateCreateCommand_Abnormal_NotExist(t *testing.T) {
	_, err := generateCreateCommand("noexists", "remote.sh")
	if err == nil {
		t.Error("Error must be occured, but it was not.")
	}
}
