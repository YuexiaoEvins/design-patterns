package facade

import "fmt"

type ComponentA struct {
	dep string
}

func NewComponentA(dep string) *ComponentA {
	return &ComponentA{
		dep: dep,
	}
}

func (a *ComponentA) Init() {
	fmt.Println("ComponentA init " + a.dep)
}

func (a *ComponentA) Operate() {
	fmt.Println("Operate A")
}
func (b *ComponentB) Operate() {
	fmt.Println("Operate A")
}

type ComponentB struct {
	dep string
}

func NewComponentB(dep string) *ComponentB {
	return &ComponentB{
		dep: dep,
	}
}

func (b *ComponentB) Init() {
	fmt.Println("ComponentB init! " + b.dep)
}
