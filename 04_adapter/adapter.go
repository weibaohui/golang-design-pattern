package adapter

import "fmt"

// Target 是 客户端使用的接口，与特定领域相关
type Target interface {
	// 示意方法，客户端请求处理的方法
	Request()
}

type TargetImpl struct {
	adaptee *Adaptee
}

func (t *TargetImpl) Request() {
	t.adaptee.SpecificRequest()
}

// Adaptee 已经存在的接口，需要被适配
type Adaptee struct{}

// 原本已经存在，已经实现过了的方法
func (a *Adaptee) SpecificRequest() {
	fmt.Println("SpecificRequest", "原本已经存在，已经实现过了的方法")
}
