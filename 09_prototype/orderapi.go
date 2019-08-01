package prototype

type OrderApi interface {
	GetOrderProductNum() int32
	SetOrderProductNum(num int32)
	CloneOrder() OrderApi
}
