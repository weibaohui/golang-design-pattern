package prototype

import (
	"fmt"
)

// 个人订单
type Personal struct {
	CustomerName    string
	ProductId       string
	OrderProductNum int32
}

func (p *Personal) SetOrderProductNum(num int32) {
	p.OrderProductNum = num
}

func (p *Personal) GetOrderProductNum() int32 {
	return p.OrderProductNum
}

func (p *Personal) CloneOrder() OrderApi {
	return &Personal{
		CustomerName:    p.CustomerName,
		ProductId:       p.ProductId,
		OrderProductNum: p.OrderProductNum,
	}
}

func (p *Personal) String() string {
	return fmt.Sprintf("本个人订单的订购人是=%s,订购产品=%s,订购数量=%d",
		p.CustomerName, p.ProductId, p.OrderProductNum)
}
