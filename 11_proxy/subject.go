package proxy

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct{}

func (s *RealSubject) Request() {
	fmt.Println("RealSubject")
}

type Proxy struct {
	RealSubject *RealSubject
}

func (s *Proxy) Request() {
	fmt.Println("before real request")
	s.RealSubject.Request()
	fmt.Println("after real request")
}
