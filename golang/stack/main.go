package main

import (
	"datastructs/golang/stack/modules"
	"fmt"
)

func main() {
	//basicStack()
	modules.RevString()
}

func basicStack() {
	fmt.Println("Initializing...")
	fmt.Println("--------------------")

	var stack modules.Stack
	fmt.Printf("Stack Adrress: %p\n", &stack)
	fmt.Println("--------------------")
	stack.IsFull()

	for i := 1; i <= modules.MaxCap; i++ {
		stack.Push(i)
	}

	stack.IsFull()
	fmt.Println("--------------------")
	fmt.Printf("Stack Adrress: %v\n", stack)
	fmt.Println("--------------------")

	for i := 1; i <= 101; i++ {
		stack.Pop()
	}

	fmt.Println("--------------------")
	fmt.Println("Done...")
}
