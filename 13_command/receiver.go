package command

import "fmt"

type Receiver struct {
}

func (r *Receiver) Action() {
	fmt.Println("Receiver 执行命令操作的功能代码")
}
