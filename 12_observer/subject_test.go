package observer

import "testing"

func TestRun(t *testing.T) {
	newspaper := NewSubject()

	r1 := &reader{Name: "张三"}
	r2 := &reader{Name: "李四"}
	r3 := &reader{Name: "王五"}
	newspaper.Attach(r1)
	newspaper.Attach(r2)
	newspaper.Attach(r3)

	newspaper.SetContent("第一期内容是观察者模式")

	newspaper.Detach(r3)
	newspaper.SetContent("第二期内容是少一个订阅")

}
