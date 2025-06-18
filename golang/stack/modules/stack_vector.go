package modules

import (
	"fmt"
)

const MaxCap = 10

type Stack struct {
	El []int
}

func (s *Stack) Push(value int) {
	if len(s.El) == MaxCap {
		fmt.Println("Max. capacity!")
		return
	}

	//fmt.Printf("-> Actual state of stack:\n, %v", s.El)
	s.El = append(s.El, value)
}

func (s *Stack) Pop() int {
	if len(s.El) == 0 {
		fmt.Println("Stack is empty!")
		return -1
	}

	lastEl := s.El[len(s.El)-1]
	s.El = s.El[:len(s.El)-1]

	//fmt.Printf("Element removed from stack: %d\n", lastEl)
	return lastEl
}

func (s *Stack) IsFull() bool {
	if len(s.El) == MaxCap {
		fmt.Printf("Stack is filled!\n")
		return true
	}

	remain := MaxCap - len(s.El)
	fmt.Printf("Remaining capacity: %d\n", remain)
	return false
}

func (s *Stack) Len() int {
	return len(s.El)
}
