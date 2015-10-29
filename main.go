package main

import (
	"errors"
	"flag"
	"os"

	"github.com/unirita/remexec/config"
	"github.com/unirita/remexec/console"
)

type arguments struct {
	configPath string
	command    string
	scriptPath string
}

const (
	rc_OK    = 0
	rc_ERROR = 1
)

func main() {
	os.Exit(realMain(fetchArgs()))
}

func realMain(args *arguments) int {
	if err := validateArgs(args); err != nil {
		console.Display("REX001E", err)
		return rc_ERROR
	}

	if err := config.Load(args.configPath); err != nil {
		console.Display("REX002E", err)
		return rc_ERROR
	}

	return rc_OK
}

func fetchArgs() *arguments {
	a := new(arguments)
	flag.StringVar(&a.configPath, "c", "", "config file path")
	flag.StringVar(&a.command, "e", "", "command to execute")
	flag.StringVar(&a.scriptPath, "f", "", "script file path to execute")
	flag.Parse()
	return a
}

func validateArgs(args *arguments) error {
	if args.configPath == "" {
		return errors.New("Set config file path with -c option.")
	}
	if args.command == "" && args.scriptPath == "" {
		return errors.New("Set command(-e) or script file path(-f).")
	}
	if args.command != "" && args.scriptPath != "" {
		return errors.New("Can not set both of -e option and -f option.")
	}
	return nil
}
