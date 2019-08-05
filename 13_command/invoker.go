package command

import "fmt"

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) RunCommand() {
	if i.command == nil {
		fmt.Println("还没有给 Invoker 设置command")
		return
	}
	i.command.Execute()
}
