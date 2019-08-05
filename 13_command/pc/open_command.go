package pc

type OpenCommand struct {
	mainboard MainBoardApi
}

func NewOpenCommand(m MainBoardApi) *OpenCommand {
	return &OpenCommand{mainboard: m}
}

func (c *OpenCommand) Execute() {
	// command 对象不知道如何开机，但可以调用主板的开机
	// 让主板去开机
	c.mainboard.Open()
}
