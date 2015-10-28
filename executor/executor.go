// Package executor provides interfaces and structs to execute remote commands or scripts.
package executor

type Executor interface {
	ExecuteCommand(string) error
	ExecuteScript(string) error
}
