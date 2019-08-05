package pc

type ResetCommand struct {
	mainboard MainBoardApi
}

func NewResetCommand(m MainBoardApi) *ResetCommand {
	return &ResetCommand{mainboard: m}
}

func (c *ResetCommand) Execute() {
	// command 对象不知道如何重启，但可以调用主板的重启
	// 让主板去重启
	c.mainboard.Reset()
}
