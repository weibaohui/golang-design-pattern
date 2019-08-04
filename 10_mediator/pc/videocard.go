package pc

import "fmt"

type VideoCard struct {
	mediator Mediator
	data     string
}

func NewVideoCard(m Mediator) *VideoCard {
	return &VideoCard{
		mediator: m,
	}
}
func (c *VideoCard) ShowData(data string) {
	fmt.Println("您正在观看的是：", data)
}
