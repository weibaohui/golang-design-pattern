package simplefactory

import "testing"

func TestInit(t *testing.T) {
	api := Init(1)
	api.Operation("简单工厂测试implA")
}

func TestInit2(t *testing.T) {
	api := Init(2)
	api.Operation("简单工厂测试implB")
}
