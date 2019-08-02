package dptuser

type UserModel struct {
	userID string
	name   string
	DepID  string
	Sex    string
}

func (m *UserModel) GetUserId() string {
	return m.userID
}

func (m *UserModel) SetUserId(id string) {
	m.userID = id
}
func (m *UserModel) GetName() string {
	return m.name
}
func (m *UserModel) SetName(name string) {
	m.name = name
}

func (m *UserModel) GetSex() string {
	return m.Sex
}
func (m *UserModel) SetSex(sex string) {
	m.Sex = sex
}

func (m *UserModel) GetDepId() string {
	return m.DepID
}
func (m *UserModel) SetDepId(id string) {
	m.DepID = id
}
