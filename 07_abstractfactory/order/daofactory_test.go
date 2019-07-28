package order

import "testing"

func TestRun(t *testing.T) {
	// 使用XML抽象工厂
	Run(&XmlDAOFactory{})
	// 使用RDB抽象工厂
	Run(&RdbDAOFactory{})
}

// 使用DAO的抽象工厂
func Run(factory DAOFactory) {
	// 通过抽象工厂获取需要的DAO接口
	mainDAO := factory.CreateOrderMainDAO()
	detailDAO := factory.CreateOrderDetailDAO()

	// 调用DAO完成存储功能
	mainDAO.SaveOrderMain()
	detailDAO.SaveOrderDetail()
}
