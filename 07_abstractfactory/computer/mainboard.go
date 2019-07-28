package computer

import "fmt"

// 主板接口
type MainboardApi interface {
	InstallCPU()
}

type GAMainboard struct {
	cpuHoles int
}

func NewGAMainboard(cpuHoles int) *GAMainboard {
	return &GAMainboard{cpuHoles: cpuHoles}
}
func (m *GAMainboard) InstallCPU() {
	fmt.Printf("now in GAMainboard,cpuHoles = %d \n", m.cpuHoles)
}

type MSIMainboard struct {
	cpuHoles int
}

func NewMSIMainboard(cpuHoles int) *MSIMainboard {
	return &MSIMainboard{cpuHoles: cpuHoles}
}
func (m *MSIMainboard) InstallCPU() {
	fmt.Printf("now in MSIMainboard,cpuHoles = %d \n", m.cpuHoles)
}
