package modules

import (
	"fmt"
)

const MaxCap = 100

type Stack struct {
	el []int
}

func (s *Stack) Push(value int) {
	if len(s.el) == MaxCap {
		fmt.Println("Max. capacity!")
		return
	}

	s.el = append(s.el, value)
}

func (s *Stack) Pop() int {
	if len(s.el) == 0 {
		fmt.Println("Stack is empty!")
		return -1
	}

	lastEl := s.el[len(s.el)-1]
	s.el = s.el[:len(s.el)-1]

	//fmt.Printf("Element removed from stack: %d\n", lastEl)
	return lastEl
}

func (s *Stack) IsFull() bool {
	if len(s.el) == MaxCap {
		fmt.Printf("Stack is filled!\n")
		return true
	}

	remain := MaxCap - len(s.el)
	fmt.Printf("Remaining capacity: %d\n", remain)
	return false
}

func (s *Stack) Len() int {
	return len(s.el)
}
