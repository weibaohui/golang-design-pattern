package singleton

import (
	"fmt"
	"testing"
)

func TestNewInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println("懒汉式", instance1, instance2, instance2 == instance1)
	if instance1 != instance2 {
		t.Fatal("懒汉式 单例失败")
	}

}

func TestGetHungryInstance(t *testing.T) {

	lazy1 := GetHungryInstance()
	lazy2 := GetHungryInstance()
	fmt.Println("饿汉式1", lazy1, lazy2, lazy1 == lazy2)
	fmt.Println("饿汉式2", hungryInstance, hungryInstance2, hungryInstance == hungryInstance2)

	if lazy1 != lazy2 {
		t.Fatal("饿汉式 单例失败")
	}
}
