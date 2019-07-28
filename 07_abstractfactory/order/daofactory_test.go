package order

import "testing"

func TestRun(t *testing.T) {
	Run(&XmlDAOFactory{})
	Run(&RdbDAOFactory{})
}
func Run(factory DAOFactory) {
	mainDAO := factory.CreateOrderMainDAO()
	detailDAO := factory.CreateOrderDetailDAO()
	mainDAO.SaveOrderMain()
	detailDAO.SaveOrderDetail()
}
