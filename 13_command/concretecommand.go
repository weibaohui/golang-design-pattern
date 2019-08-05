package command

type ConcreteCommand struct {
	state    string
	receiver *Receiver
}

func NewConcreteCommand(r *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		state:    "",
		receiver: r,
	}
}
func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}
