package pc

import "fmt"

type GigaMainBoard struct {
}

func (m *GigaMainBoard) Open() {
	fmt.Println("技嘉主板开机")
	fmt.Println("接通电源")
	fmt.Println("设备检查")
	fmt.Println("装载系统")
	fmt.Println("正常运转")
	fmt.Println("已开机")
}

func (m *GigaMainBoard) Reset() {
	fmt.Println("技嘉主板正在重启")
	fmt.Println("已开机")
}
