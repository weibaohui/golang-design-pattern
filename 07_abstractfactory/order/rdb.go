package order

import "fmt"

type RdbMainDAOImpl struct{}
type RdbDetailDAOImpl struct{}

func (m *RdbMainDAOImpl) SaveOrderMain() {
	fmt.Printf("now in RdbMainDAOImpl SaveOrderMain \n")
}

func (m *RdbDetailDAOImpl) SaveOrderDetail() {
	fmt.Printf("now in RdbMainDAOImpl SaveOrderDetail \n")
}

type RdbDAOFactory struct{}

func (d *RdbDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RdbMainDAOImpl{}
}

func (d *RdbDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RdbDetailDAOImpl{}
}
