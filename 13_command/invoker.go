package command

type Invoker struct {
	command Command
}

func NewInvoker(c Command) *Invoker {
	return &Invoker{command: c}
}
func (i *Invoker) RunCommand() {
	i.command.Execute()
}
