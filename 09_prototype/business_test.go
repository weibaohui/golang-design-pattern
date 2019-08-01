package prototype

import (
	"testing"
)

func TestPersonalBusiness(t *testing.T) {
	personalOrder := &Personal{
		CustomerName:    "个人A",
		ProductId:       "10000A",
		OrderProductNum: 1001,
	}
	b := OrderBusiness{}
	b.SaveOrder(personalOrder)
}

func TestEnterpriseBusiness(t *testing.T) {
	enterpriseOrder := &Enterprise{
		EnterpriseName:  "企业C",
		ProductId:       "Ex00000AC",
		OrderProductNum: 1002,
	}
	b := OrderBusiness{}
	b.SaveOrder(enterpriseOrder)
}
