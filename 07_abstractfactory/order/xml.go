package order

import "fmt"

type XmlMainDAOImpl struct{}
type XmlDetailDAOImpl struct{}

func (m *XmlMainDAOImpl) SaveOrderMain() {
	fmt.Printf("now in XmlMainDAOImpl SaveOrderMain \n")
}

func (m *XmlDetailDAOImpl) SaveOrderDetail() {
	fmt.Printf("now in XmlMainDAOImpl SaveOrderDetail \n")
}

type XmlDAOFactory struct{}

func (d *XmlDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XmlMainDAOImpl{}
}

func (d *XmlDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XmlDetailDAOImpl{}
}
