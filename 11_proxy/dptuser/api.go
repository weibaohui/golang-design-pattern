package dptuser

type UserModelApi interface {
	GetUserId() string
	SetUserId(id string)
	GetName() string
	SetName(name string)
	GetDepId() string
	SetDepId(id string)
	GetSex() string
	SetSex(sex string)
}
