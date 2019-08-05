package command

import "testing"

func TestRun(t *testing.T) {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := NewInvoker(command)
	invoker.RunCommand()
}
