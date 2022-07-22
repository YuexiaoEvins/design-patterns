package creational

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := NewSingleton()
	fmt.Printf("inst: %s addr %v ori: %v\n", instance1, &instance1, instance1)
	instance2 := NewSingleton()
	fmt.Printf("inst: %s addr %v ori: %v\n", instance2, &instance2, instance2)
}
