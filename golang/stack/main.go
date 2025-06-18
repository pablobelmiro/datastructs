package main

import (
	"datastructs/golang/stack/modules"
	"fmt"
	"strconv"
)

func main() {
	s := modules.Stack{
		El: make([]int, 0, 10),
	}

	for {
		fmt.Println("---------------------")
		var input string
		fmt.Println("Enter a stack function:")
		fmt.Printf(
			"1. Manipulate stack\n" +
				"2. Reverse word using stack\n" +
				"3. Undo integer Stack\n" +
				"4. Stop\n")
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Only number inputs to enter a function!")
			continue
		}

		switch num {
		case 1:
			basicStack(&s)
		case 2:
			reverseString(&s)
		case 3:
			undoA(&s)
		case 4:
			return
		}
	}

}

func basicStack(s *modules.Stack) {
	fmt.Println("Stack manipulation...")

	for {
		fmt.Println("-------")
		var input string
		fmt.Printf(
			"1. Add number\n" +
				"2. Pop last item\n" +
				"3. Check Length\n" +
				"4. Check if is filled\n" +
				"5. Stop\n")
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Only number inputs to enter a function!")
			continue
		}

		switch num {
		case 1:
			pushVal(s)
		case 2:
			popVal(s)
		case 3:
			checkLen(s)
		case 4:
			s.IsFull()
		case 5:
			return
		}
	}
}

func pushVal(s *modules.Stack) {
	fmt.Println("Enter a number:")
	var input string
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Only number can be insert!")
		return
	}

	s.Push(num)
}

func popVal(s *modules.Stack) {
	s.Pop()
}

func checkLen(s *modules.Stack) {
	fmt.Printf("Len: %d\n", len(s.El))
	fmt.Printf("Cap: %d\n", cap(s.El))
}

func reverseString(s *modules.Stack) {
	modules.RevString(s)
}

func checkAddr(s *modules.Stack) {
	fmt.Printf("Stack Adrress: %p\n", &s)
	fmt.Println("_______________")
}

func undoA(s *modules.Stack) {
	modules.UndoAction(s)
}
