package template

type NormalLogin struct {
	LoginTemplate
}

func (lt *NormalLogin) FindLoginUser(loginID string) *LoginModel {
	return &LoginModel{
		LoginID:  loginID,
		Password: "normal_password",
	}
}
