package prototype

import (
	"fmt"
)

// 企业订单
type Enterprise struct {
	EnterpriseName  string
	ProductId       string
	OrderProductNum int32
}

func (p *Enterprise) SetOrderProductNum(num int32) {
	p.OrderProductNum = num
}

func (p *Enterprise) GetOrderProductNum() int32 {
	return p.OrderProductNum
}

func (p *Enterprise) CloneOrder() OrderApi {
	return &Enterprise{
		EnterpriseName:  p.EnterpriseName,
		ProductId:       p.ProductId,
		OrderProductNum: p.OrderProductNum,
	}
}

func (p *Enterprise) String() string {
	return fmt.Sprintf("本企业订单的订购企业是=%s,订购产品=%s,订购数量=%d",
		p.EnterpriseName, p.ProductId, p.OrderProductNum)
}
