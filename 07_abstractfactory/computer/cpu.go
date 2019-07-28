package computer

import "fmt"

// cpu 接口
type CPUApi interface {
	Calculate()
}

type IntelCpu struct {
	pins int // cpu针脚数量
}

func NewIntelCpu(pins int) *IntelCpu {
	return &IntelCpu{pins: pins}
}
func (c *IntelCpu) Calculate() {
	fmt.Printf("now in Intel CPU,pins= %d \n", c.pins)
}

type AMDCpu struct {
	pins int // cpu针脚数量
}

func NewAMDCpu(pins int) *AMDCpu {
	return &AMDCpu{pins: pins}
}
func (c *AMDCpu) Calculate() {
	fmt.Printf("now in AMD CPU,pins= %d \n", c.pins)
}
