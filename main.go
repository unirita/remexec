package main

import (
	"fmt"
	"os"

	"github.com/unirita/remexec/executor"
)

//リターンコード
var (
	RC_OK    = 0
	RC_ERROR = 1
)

func main() {
	os.Exit(realMain())
}

func realMain() int {

	exec := executor.New("testserver", "Administrator", "Adm1n", "windows")

	if err := exec.ExecuteCommand("dir c:\\"); err != nil {
		fmt.Println(err)
		return RC_ERROR
	}

	return RC_OK
}
