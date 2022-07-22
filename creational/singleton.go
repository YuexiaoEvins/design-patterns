package creational

import "sync"

type Singleton struct {
	name string
}

func (s *Singleton) String() string {
	return s.name
}

func NewSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{name: "test"}
	})
	return instance
}

var (
	once     sync.Once
	instance *Singleton
)
