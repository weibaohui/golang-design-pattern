package pc

import command "github.com/weibaohui/design-pattern/13_command"

type Box struct {
	openCommand  command.Command
	resetCommand command.Command
}

func (b *Box) SetOpenCommand(c command.Command) {
	b.openCommand = c
}

func (b *Box) SetResetCommand(c command.Command) {
	b.resetCommand = c
}

func (b *Box) ResetPressed() {
	b.resetCommand.Execute()
}

func (b *Box) OpenPressed() {
	b.openCommand.Execute()
}
