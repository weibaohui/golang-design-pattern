package template

import "fmt"

type Loginer interface {
	FindLoginUser(loginID string) *LoginModel
	EncryptPwd(pwd string) string
	Match(model *LoginModel, dblm *LoginModel) bool
}

type LoginTemplate struct {
	impl Loginer
}

func NewLoginTemplate(loginer Loginer) *LoginTemplate {
	return &LoginTemplate{impl: loginer}
}

// 登录是否成功
func (lt *LoginTemplate) Login(lm *LoginModel) bool {
	// 1 根据登录人员的编号去获取相应的数据
	user := lt.impl.FindLoginUser(lm.LoginID)
	if user != nil {
		// 2 对密码进行加密,并更新到lm数据模型
		pwd := lt.impl.EncryptPwd(lm.Password)
		lm.Password = pwd
		// 3 判断是否匹配
		return lt.impl.Match(lm, user)
	}

	return false
}
func (lt *LoginTemplate) FindLoginUser(loginID string) *LoginModel {
	fmt.Println("未实现")
	return nil
}
func (lt *LoginTemplate) EncryptPwd(pwd string) string {
	return pwd
}

// 判断 用户名密码 是否跟数据库中的匹配
func (lt *LoginTemplate) Match(model *LoginModel, dblm *LoginModel) bool {
	return model.Password == dblm.Password && model.LoginID == dblm.LoginID
}
