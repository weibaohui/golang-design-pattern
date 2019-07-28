package order

// 抽象工厂，创建主订单、订单详情的DAO对象
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// 主订单
type OrderMainDAO interface {
	// 保存主订单
	SaveOrderMain()
}

// 订单详情
type OrderDetailDAO interface {
	// 保存订单详情
	SaveOrderDetail()
}
