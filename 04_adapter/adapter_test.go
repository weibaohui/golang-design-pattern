package adapter

import "testing"

func TestTargetImpl_Request(t *testing.T) {
	target := TargetImpl{adaptee: &Adaptee{}}
	target.Request()
}
