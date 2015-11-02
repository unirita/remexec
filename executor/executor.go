// Package executor provides interfaces and structs to execute remote commands or scripts.
package executor

type Executor interface {
	ExecuteCommand(string) (int, error)
	ExecuteScript(string) (int, error)
}
