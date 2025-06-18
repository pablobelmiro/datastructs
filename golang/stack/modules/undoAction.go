package modules

import (
	"fmt"
	"strconv"
)

func UndoAction(s *Stack) {
	fmt.Println("Starting...")
	s.El = s.El[:0]

	for {
		var input string
		fmt.Println("Insert a number(back to undo)")
		fmt.Scan(&input)

		if input == "stop" {
			return
		}

		if input == "back" {
			if len(s.El) == 0 {
				fmt.Println("List is empty!")
				continue
			}
			removed := s.Pop()
			if len(s.El) == 0 {
				fmt.Println("List is empty!")
				continue
			}
			fmt.Println("-----------")
			fmt.Printf("Back from: %d\n", removed)
			fmt.Printf("Actual 'element:' %d\n", s.El[len(s.El)-1])
			fmt.Printf("stack:' %v\n", s.El)
			fmt.Println("-----------")
			continue
		}

		InConverted, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Input must be a integer!")
			continue
		}

		addEl(s, InConverted)
	}
}

func addEl(s *Stack, value int) {
	s.Push(value)
	fmt.Printf("Actual Element: %d\n", s.El[len(s.El)-1])
}
