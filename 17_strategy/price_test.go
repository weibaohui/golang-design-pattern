package strategy

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	strategy := &LargeCustomerStrategy{}
	price := NewPrice(strategy)
	result := price.quote(7876)
	fmt.Println("报价", result)
}
