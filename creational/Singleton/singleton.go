package Singleton

import "sync"

type Singleton struct {
	name string
	val  int
}

func NewSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{name: "test", val: 1}
	})
	return instance
}

var (
	once     sync.Once
	instance *Singleton
)
