package strategy

type Price struct {
	Strategy Strategy
}

func NewPrice(s Strategy) *Price {
	return &Price{Strategy: s}
}
func (c *Price) quote(price float64) float64 {
	return c.Strategy.CalcPrice(price)
}
