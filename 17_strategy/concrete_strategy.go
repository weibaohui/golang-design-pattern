package strategy

import "fmt"

type NormalCustomerStrategy struct {
}

func (c *NormalCustomerStrategy) CalcPrice(price float64) float64 {
	fmt.Println("普通客户，没有折扣")
	return price
}

type OldCustomerStrategy struct {
}

func (c *OldCustomerStrategy) CalcPrice(price float64) float64 {
	fmt.Println("老客户，折扣5%")
	return price * (1 - 0.05)
}

type LargeCustomerStrategy struct {
}

func (c *LargeCustomerStrategy) CalcPrice(price float64) float64 {
	fmt.Println("大客户，折扣10%")
	return price * (1 - 0.1)
}
