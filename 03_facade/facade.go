package facade

import "fmt"

func Facade() {
	api := &ApiImpl{
		a: &AModuleImpl{},
		b: &BModuleImpl{},
		c: &CModuleImpl{},
	}
	api.Test()
}

type Api interface {
	Test()
}
type ApiImpl struct {
	a AModuleApi
	b BModuleApi
	c CModuleApi
}

func (i *ApiImpl) Test() {
	i.a.testA()
	i.b.testB()
	i.c.testC()
}

type AModuleApi interface {
	testA()
}

type AModuleImpl struct{}

func (i *AModuleImpl) testA() {
	fmt.Println("现在在A模块里面操作TestA方法")
}

type BModuleApi interface {
	testB()
}

type BModuleImpl struct{}

func (i *BModuleImpl) testB() {
	fmt.Println("现在在B模块里面操作TestB方法")
}

type CModuleApi interface {
	testC()
}

type CModuleImpl struct{}

func (i *CModuleImpl) testC() {
	fmt.Println("现在在C模块里面操作TestC方法")
}
