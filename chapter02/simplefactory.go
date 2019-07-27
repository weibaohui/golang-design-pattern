package chapter02

import "fmt"

type Api interface {
	Operation(s string)
}

func Init(condition int) Api {
	if condition == 1 {
		return &implA{}
	} else if condition == 2 {
		return &implB{}
	}
	return nil
}

type implA struct{}

func (i *implA) Operation(s string) {
	fmt.Println("ImplA s==", s)
}

type implB struct{}

func (i *implB) Operation(s string) {
	fmt.Println("ImplB s==", s)
}
