package dptuser

type Manager struct{}

func (m *Manager) GetUserByDepId() []UserModelApi {
	models := make([]UserModelApi, 0)
	proxy1 := &Proxy{RealSubject: &UserModel{}}
	proxy1.SetName("张三丰")
	proxy1.SetUserId("zhangsf")
	models = append(models, proxy1)

	proxy2 := &Proxy{RealSubject: &UserModel{}}
	proxy2.SetName("李四")
	proxy2.SetUserId("lis")
	models = append(models, proxy2)

	return models
}
