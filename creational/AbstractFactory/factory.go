package AbstractFactory

import "fmt"

type ObjectA interface {
	OperateA()
}

type ObjectB interface {
	OperateB()
}

type AImpl1 struct {
}

func (a AImpl1) OperateA() {
	fmt.Println("AImpl1 OperateA")
}

type AImpl2 struct {
}

func (a AImpl2) OperateA() {
	fmt.Println("AImpl2 OperateA")
}

type BImpl1 struct {
}

func (b BImpl1) OperateB() {
	fmt.Println("BImpl1 OperateB")
}

type BImpl2 struct {
}

func (b BImpl2) OperateB() {
	fmt.Println("BImpl2 OperateB")
}

type Creator1 struct {
}

func (c *Creator1) CreateObjectA() ObjectA {
	return &AImpl1{}
}
func (c *Creator1) CreateObjectB() ObjectB {
	return &BImpl1{}
}

type Creator2 struct {
}

func (c *Creator2) CreateObjectA() ObjectA {
	return &AImpl2{}
}
func (c *Creator2) CreateObjectB() ObjectB {
	return &BImpl2{}
}

type AbstractFactory interface {
	CreateObjectA() ObjectA
	CreateObjectB() ObjectB
}
