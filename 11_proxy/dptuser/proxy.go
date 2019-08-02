package dptuser

import "fmt"

type Proxy struct {
	RealSubject *UserModel
	loaded      bool
}

func (m *Proxy) GetUserId() string {
	return m.RealSubject.GetUserId()
}

func (m *Proxy) SetUserId(id string) {
	m.RealSubject.SetUserId(id)
}
func (m *Proxy) GetName() string {
	return m.RealSubject.GetName()
}
func (m *Proxy) SetName(name string) {
	m.RealSubject.SetName(name)
}

func (m *Proxy) GetSex() string {
	if !m.loaded {
		m.reload()
		m.loaded = true
	}
	return m.RealSubject.GetSex()
}
func (m *Proxy) SetSex(sex string) {
	m.RealSubject.SetSex(sex)
}

func (m *Proxy) GetDepId() string {
	if !m.loaded {
		m.reload()
		m.loaded = true
	}
	return m.RealSubject.GetDepId()
}

func (m *Proxy) SetDepId(id string) {
	m.RealSubject.SetDepId(id)
}

func (m *Proxy) reload() {
	fmt.Println("加载完整数据 for ", m.RealSubject.GetName())
	m.RealSubject.SetSex("男")
	m.RealSubject.SetDepId("A10010")
}
