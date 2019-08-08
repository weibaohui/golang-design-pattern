package template

import "fmt"

type WorkerLogin struct {
	LoginTemplate
}

func (lt *WorkerLogin) FindLoginUser(loginID string) *LoginModel {
	return &LoginModel{
		LoginID:  loginID,
		Password: "worker_password",
	}
}
func (lt *WorkerLogin) EncryptPwd(pwd string) string {
	fmt.Println("MD5 加密 密码")
	return pwd
}
