package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/unirita/remexec/config"
)

type WinRMExecutor struct {
	host string
	user string
	pass string
}

const (
	WINRM_CMD   = "e"
	WINRM_LOCAL = "f"
)

var remexecPs1 = filepath.Join(filepath.Dir(getModulePath()), "remexec.ps1")

type commandRunFunc func(*exec.Cmd) error

var cmdRun commandRunFunc = run

func NewWinRMExecutor(cfg *config.Config) *WinRMExecutor {
	e := new(WinRMExecutor)
	e.host = cfg.Remote.Host
	e.user = cfg.Remote.User
	e.pass = cfg.Remote.Pass

	return e
}

func (e *WinRMExecutor) ExecuteCommand(command string) (int, error) {
	cmd := e.createCmd(command, WINRM_CMD)
	return e.executeWinRM(cmd)
}

func (e *WinRMExecutor) ExecuteScript(path string) (int, error) {
	cmd := e.createCmd(path, WINRM_LOCAL)
	return e.executeWinRM(cmd)
}

func (e *WinRMExecutor) executeWinRM(cmd *exec.Cmd) (int, error) {
	rc, err := e.getRC(cmdRun(cmd))
	if err != nil {
		return -1, fmt.Errorf("Run command error: %s", err)
	}

	return rc, nil
}

func (e *WinRMExecutor) createCmd(execution string, scripttype string) *exec.Cmd {
	cmd := new(exec.Cmd)
	powershell := filepath.Join(os.Getenv("PSModulePath"), "..", "powershell.exe")
	option := "-File"

	cmd = exec.Command(powershell, option, remexecPs1, scripttype, e.host, e.user, e.pass, execution)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func run(cmd *exec.Cmd) error {
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (e *WinRMExecutor) getRC(err error) (int, error) {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
			panic(errors.New("Unimplemented for system where exec.ExitError.Sys() is not syscall.WaitStatus."))
		}
		return -1, err
	}
	return 0, nil
}

// DLLハンドル
type cutoDLL struct {
	dll *syscall.DLL
}

var (
	kernel32_dll           = loadDLL("kernel32.dll")
	procGetModuleFileNameW = kernel32_dll.findProc("GetModuleFileNameW")
)

func getModulePath() string {
	// MAX_PATHがUTF-8になる場合は、これくらいあれば十分か？
	const max_path = 520
	var buf [max_path]byte
	procGetModuleFileNameW.Call(0, uintptr(unsafe.Pointer(&buf)), (uintptr)(max_path))

	// Unicodeで取得しているので、2byte目が0の部分を除外する。
	var path [max_path / 2]byte
	var j int
	for i := 0; i < len(buf); i++ {
		if buf[i] != 0 {
			path[j] = buf[i]
			j++
		}
	}
	return fmt.Sprintf("%s", path)
}

func loadDLL(name string) *cutoDLL {
	dll, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	cutoDll := new(cutoDLL)
	cutoDll.dll = dll
	return cutoDll
}

func (c *cutoDLL) findProc(name string) *syscall.Proc {
	proc, err := c.dll.FindProc(name)
	if err != nil {
		panic(err)
	}
	return proc
}
