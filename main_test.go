package main

import (
	"flag"
	"os"
	"testing"
)

func TestFetchArgs(t *testing.T) {
	os.Args = os.Args[:1]
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append(os.Args,
		"-c", "test.ini", "-e", "testcommand", "-f", "test.sh")

	args := fetchArgs()
	if args.configPath != "test.ini" {
		t.Errorf("args.configPath => %s, wants %s", args.configPath, "test.ini")
	}
	if args.command != "testcommand" {
		t.Errorf("args.command => %s, wants %s", args.command, "testcommand")
	}
	if args.scriptPath != "test.sh" {
		t.Errorf("args.scriptPath => %s, wants %s", args.scriptPath, "test.sh")
	}
}

func TestFetchArgs_NoOption(t *testing.T) {
	os.Args = os.Args[:1]
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	args := fetchArgs()
	if args.configPath != "" {
		t.Errorf("args.configPath => %s, must be empty.", args.configPath)
	}
	if args.command != "" {
		t.Errorf("args.command => %s, must be empty.", args.command)
	}
	if args.scriptPath != "" {
		t.Errorf("args.scriptPath => %s, must be empty.", args.scriptPath)
	}
}

func TestValidateArgs_Normal_WithCommand(t *testing.T) {
	args := new(arguments)
	args.configPath = "test.ini"
	args.command = "testcommand"

	err := validateArgs(args)
	if err != nil {
		t.Errorf("Error was returned unexpectedly.")
		t.Logf("Error: %s", err)
	}
}

func TestValidateArgs_Normal_WithScriptPath(t *testing.T) {
	args := new(arguments)
	args.configPath = "test.ini"
	args.scriptPath = "test.sh"

	err := validateArgs(args)
	if err != nil {
		t.Errorf("Error was returned unexpectedly.")
		t.Logf("Error: %s", err)
	}
}

func TestValidateArgs_Abnormal_NoConfigPath(t *testing.T) {
	args := new(arguments)
	args.command = "testcommand"

	err := validateArgs(args)
	if err == nil {
		t.Errorf("Error must be returned, but it was not.")
	}
}

func TestValidateArgs_Abnormal_OnlyConfigPath(t *testing.T) {
	args := new(arguments)
	args.configPath = "test.ini"

	err := validateArgs(args)
	if err == nil {
		t.Errorf("Error must be returned, but it was not.")
	}
}

func TestValidateArgs_Abnormal_BothCommandAndScript(t *testing.T) {
	args := new(arguments)
	args.configPath = "test.ini"
	args.command = "testcommand"
	args.scriptPath = "test.sh"

	err := validateArgs(args)
	if err == nil {
		t.Errorf("Error must be returned, but it was not.")
	}
}
