package facade

import "testing"

func TestFacade(t *testing.T) {
	api := Init()
	api.Test()
}
