// Describe the test of remexec.ps1.

package executor

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func makeLocalHostWinrmExecutor() *WinRMExecutor {
	e := new(WinRMExecutor)
	e.host = "localhost"
	e.user = ""
	e.pass = ""

	return e
}

func TestRemexecPs1_Command(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	cmd.Stderr = nil
	cmd.Stdout = nil

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 0 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 0)
	}
}

func TestRemexecPs1_CommandWithArgumentSuccess(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("dir c:\\", WINRM_CMD)

	cmd.Stderr = nil
	cmd.Stdout = nil

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 0 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 0)
	}
}

func TestRemexecPs1_CommandWithArgumentFailed(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("dir noexist", WINRM_CMD)

	cmd.Stderr = nil
	cmd.Stdout = nil

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 255 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 0)
	}
}

func TestRemexecPs1_RemoteBatFile(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.bat")

	cmd := e.createCmd(execution, WINRM_CMD)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 1 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 1)
	}

	if !strings.Contains(out.String(), "test message.") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_RemoteBatFileWithArgument(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.bat Arg1 Arg2 Arg3")

	cmd := e.createCmd(execution, WINRM_CMD)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 1 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 1)
	}

	if !strings.Contains(out.String(), "Arg1 Arg2 Arg3") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_RemotePSScript(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.ps1")

	cmd := e.createCmd(execution, WINRM_CMD)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 5 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 1)
	}

	if !strings.Contains(out.String(), "test message.") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_RemotePSScriptWithArgument(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.ps1 Arg1 Arg2 Arg3")

	cmd := e.createCmd(execution, WINRM_CMD)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 5 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 1)
	}

	if !strings.Contains(out.String(), "Arg1 Arg2 Arg3") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_LocalPSScript(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.ps1")

	cmd := e.createCmd(execution, WINRM_LOCAL)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	_, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if !strings.Contains(out.String(), "test message.") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_LocalPSScriptWithArgument(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.ps1 Arg1 Arg2 Arg3")

	cmd := e.createCmd(execution, WINRM_LOCAL)

	out := &bytes.Buffer{}
	cmd.Stdout = out

	_, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if !strings.Contains(out.String(), "Arg1 Arg2 Arg3") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_LocalBatFile(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	execution := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "executor", "_testdata", "remexecps1test.bat")

	cmd := e.createCmd(execution, WINRM_LOCAL)
	cmd.Stderr = nil

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 255 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 255)
	}
}

func TestRemexecPs1_InvalidArgument(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	powershell := filepath.Join(os.Getenv("PSModulePath"), "..", "powershell.exe")
	option := "-File"

	cmd := exec.Command(powershell, option, remexecPs1, WINRM_CMD, "user", "pass", "ipconfig")

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 255 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 255)
	}

	if !strings.Contains(out.String(), "Invalid argument.") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_NotExistOption(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")

	cmd := e.createCmd("ipconfig", "notexist")

	out := &bytes.Buffer{}
	cmd.Stdout = out

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 255 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 255)
	}

	if !strings.Contains(out.String(), "Unkown option notexist") {
		t.Errorf("output different. output =>[%s]", out)
	}
}

func TestRemexecPs1_NoExistCredential(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	e.host = "noexist"
	e.pass = "misspass"
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	cmd.Stderr = nil

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("An error has occurred that is not expected. %s", err)
	}

	if rc != 255 {
		t.Errorf("return code => %d, wants => %d ", rc, 255)
	}
}
