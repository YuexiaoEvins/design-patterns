package Singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := NewSingleton()
	fmt.Printf("inst: %v addr %v \n", instance1, &instance1)
	instance2 := NewSingleton()
	fmt.Printf("inst: %v addr %v \n", instance2, &instance2)
	instance1.val++
	fmt.Printf("inst: %v addr %v \n", instance2, &instance2)
	fmt.Printf("inst: %v addr %v \n", instance1, &instance1)

}
