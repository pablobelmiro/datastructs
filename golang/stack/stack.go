package main

import (
	"fmt"
)

const maxCap = 100

type Stack struct {
	el []int
}

func (s *Stack) push(value int) {
	if len(s.el) == maxCap {
		fmt.Println("Max. capacity!")
		return
	}

	s.el = append(s.el, value)
}

func (s *Stack) pop() int {
	if len(s.el) == 0 {
		fmt.Println("Stack is empty!")
		return -1
	}

	lastEl := s.el[len(s.el)-1]
	s.el = s.el[:len(s.el)-1]

	fmt.Printf("Element removed from stack: %d\n", lastEl)
	return lastEl
}

func (s *Stack) isFull() bool {
	if len(s.el) == maxCap {
		fmt.Printf("Stack is filled!\n")
		return true
	}

	remain := maxCap - len(s.el)
	fmt.Printf("Remaining capacity: %d\n", remain)
	return false
}

func main() {
	fmt.Println("Initializing...")
	fmt.Println("--------------------")

	var stack Stack
	fmt.Printf("Stack Adrress: %p\n", &stack)
	fmt.Println("--------------------")
	stack.isFull()

	for i := 1; i <= maxCap; i++ {
		stack.push(i)
	}

	stack.isFull()
	fmt.Println("--------------------")
	fmt.Printf("Stack Adrress: %v\n", stack)
	fmt.Println("--------------------")

	for i := 1; i <= 101; i++ {
		stack.pop()
	}

	fmt.Println("--------------------")
	fmt.Println("Done...")
}
