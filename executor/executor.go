package executor

type Executor interface {
	ExecuteCommand(string) error
	ExecuteScript(string) error
}
