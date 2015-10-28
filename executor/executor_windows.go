package executor

import (
	"fmt"
	"os/exec"
)

type Windows struct {
	HostName string
	UserName string
	Password string
}

type Linux struct {
}

func (w *Windows) ExecuteCommand(cm string) error {
	//TODO: Get absolute path of powershell script file.
	cmd := exec.Command("remote.ps1", "-e", "host", "user", "pass", cm)
	result, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func (w *Windows) ExecuteScript(sc string) error {
	return nil
}

func (l *Linux) ExecuteCommand(cm string) error {
	return nil
}

func (l *Linux) ExecuteScript(sc string) error {
	return nil
}

func New(host, user, pass, os string) Executor {
	// TODO: Decide struct type by os, And create struct object.
	var exec Executor

	if os == "windows" {
		exec = &Windows{HostName: host, UserName: user, Password: pass}
	} else if os == "linux" {
		exec = &Linux{}
	}
	return exec
}
