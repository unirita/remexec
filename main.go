package main

import (
	"os"
)

//return code
var (
	RC_OK    = 0
	RC_ERROR = 1
)

func main() {
	os.Exit(realMain())
}

func realMain() int {

	return RC_OK
}
