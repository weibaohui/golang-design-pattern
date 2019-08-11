package strategy

// 策略接口
type Strategy interface {
	// 某个独立算法接口，参数可根据实际情况来定义
	CalcPrice(price float64) float64
}
