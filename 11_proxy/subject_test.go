package proxy

import "testing"

func TestRun(t *testing.T) {
	realSubject := &RealSubject{}
	proxy := &Proxy{RealSubject: realSubject}
	proxy.Request()
}
