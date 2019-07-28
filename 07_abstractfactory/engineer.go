package abstractfactory

type ComputerEngineer struct {
	cpu       CPUApi
	mainboard MainboardApi
}

// 装机过程
func (c *ComputerEngineer) MakeComputer(schema AbstractFactory) {
	c.prepareHardwares(schema)
}

// 准备所需的硬件
func (c *ComputerEngineer) prepareHardwares(schema AbstractFactory) {
	// 使用抽象工厂获取相应的接口对象
	c.cpu = schema.CreateCPUApi()
	c.mainboard = schema.CreateMainboardApi()

	// 测试下硬件是否可用
	c.cpu.Calculate()
	c.mainboard.InstallCPU()
}
