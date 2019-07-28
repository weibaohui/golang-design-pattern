package singleton

import (
	"math/rand"
	"sync"
	"time"
)

type Singleton struct {
	id int64 // 随机值，验证是否为同一个实例
}

var singleton *Singleton
var once sync.Once

// 懒汉式单例
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			id: rand.NewSource(time.Now().UnixNano()).Int63(),
		}
	})
	return singleton
}

// 饿汉式单例
var hungryInstance = &Singleton{id: rand.NewSource(time.Now().UnixNano()).Int63()}
var hungryInstance2 = &Singleton{id: rand.NewSource(time.Now().UnixNano()).Int63()}

func GetHungryInstance() *Singleton {
	return hungryInstance
}
