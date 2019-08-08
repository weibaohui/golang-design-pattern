package template

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	user := &LoginModel{
		LoginID:  "admin",
		Password: "worker_password",
	}

	workerLogin := NewLoginTemplate(&WorkerLogin{})
	normalLogin := NewLoginTemplate(&NormalLogin{})

	if ok := workerLogin.Login(user); ok {
		fmt.Println("登录 工作 平台成功")
	} else {
		fmt.Println("登录 工作 平台失败")
	}
	if ok := normalLogin.Login(user); ok {
		fmt.Println("登录  普通 平台成功")
	} else {
		fmt.Println("登录  普通 平台失败")
	}
}
