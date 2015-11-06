// Describe the test of remexec.ps1.

package executor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemexecPs1_InvalidArgument(t *testing.T) {

}

func TestRemexecPs1_NoExistCredential(t *testing.T) {
	e := makeTestWinrmExecutor()
	e.host = "noexist"
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
