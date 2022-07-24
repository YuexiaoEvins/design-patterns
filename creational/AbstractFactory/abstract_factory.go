package AbstractFactory

import "fmt"

type ProductA interface {
	A()
}

type ProductB interface {
	B()
}

type FactoryImpl1 struct {
}

type ProductA1 struct {
	name string
}

func (a ProductA1) A() {
	fmt.Println("A1 A() name %s", a.name)
}

func (f FactoryImpl1) CreateProductA() ProductA {
	return &ProductA1{
		name: "A1",
	}
}

type WidgetFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}
