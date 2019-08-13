package strategy

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	strategy := &LargeCustomerStrategy{}
	price := NewPrice(strategy)
	result := price.quote(7876)

	price = NewPrice(&OldCustomerStrategy{})
	result = price.quote(7655.32)
	fmt.Println("报价", result)
}
