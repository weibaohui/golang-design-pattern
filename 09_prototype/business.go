package prototype

import "fmt"

type OrderBusiness struct{}

func (b *OrderBusiness) SaveOrder(order OrderApi) {
	for order.GetOrderProductNum() > 1000 {
		newOrder := order.CloneOrder()
		newOrder.SetOrderProductNum(1000)

		order.SetOrderProductNum(order.GetOrderProductNum() - 1000)
		fmt.Println("拆分订单==", newOrder)
	}

	fmt.Println("订单==", order)
}
