package FactoryMethod

import "fmt"

func GetProduct(pType string) Product {
	switch pType {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		fmt.Errorf("unknown type")
		return nil
	}
}

type Product interface {
	Operate()
}

type ProductA struct {
	size int
}

type ProductB struct {
	size int
}

func (a ProductA) Operate() {
	fmt.Println("Product A Operate()!")
}

func (a ProductB) Operate() {
	fmt.Println("Product B Operate()!")
}
