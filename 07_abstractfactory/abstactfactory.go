package abstractfactory

type AbstractFactory interface {
	CreateCPUApi() CPUApi
	CreateMainboardApi() MainboardApi
}

type Schema1 struct{}

func (s *Schema1) CreateCPUApi() CPUApi {
	return &IntelCpu{pins: 1156}
}

func (s *Schema1) CreateMainboardApi() MainboardApi {
	return &GAMainboard{cpuHoles: 1156}
}

type Schema2 struct{}

func (s *Schema2) CreateCPUApi() CPUApi {
	return &AMDCpu{pins: 939}
}

func (s *Schema2) CreateMainboardApi() MainboardApi {
	return &MSIMainboard{cpuHoles: 939}
}
