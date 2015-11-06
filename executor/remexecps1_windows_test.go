// Describe the test of remexec.ps1.

package executor

import (
	"os"
	"path/filepath"
	"testing"
)

func makeLocalHostWinrmExecutor() *WinRMExecutor {
	e := new(WinRMExecutor)
	e.host = "localhost"
	e.user = ""
	e.pass = ""

	return e
}

func TestRemexecPs1_InvalidArgument(t *testing.T) {
	e := makeLocalHostWinrmExecutor()
	remexecPs1 = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita", "remexec", "script", "remexec.ps1")
	cmd := e.createCmd("ipconfig", WINRM_CMD)

	cmd.Stderr = nil

	rc, err := e.runAndGetRC(cmd)

	if err != nil {
		t.Fatalf("error occurred %s", err)
	}

	if rc != 0 {
		t.Errorf("return code not match. rc => %d, expect => %d", rc, 0)
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
