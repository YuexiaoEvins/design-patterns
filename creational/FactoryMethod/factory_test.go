package FactoryMethod

import "testing"

func TestFactory(t *testing.T) {
	p := GetProduct("A")
	p.Operate()
}
