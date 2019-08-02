package dptuser

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	manager := &Manager{}
	models := manager.GetUserByDepId()
	for _, m := range models {
		fmt.Println(m.GetName(), m.GetUserId())
	}

	for _, m := range models {
		fmt.Println(m.GetName(), m.GetUserId(), m.GetSex(), m.GetDepId())
	}

	for _, m := range models {
		fmt.Println(m.GetName(), m.GetUserId(), m.GetSex(), m.GetDepId())
	}
}
