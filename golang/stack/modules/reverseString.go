package modules

import (
	"fmt"
)

func RevString() string {
	var word string
	var s Stack
	var revString string
	fmt.Println("Enter a string:")
	fmt.Scan(&word)

	for _, c := range word {
		s.Push(int(c))
	}

	for s.Len() > 0 {
		revString += string(s.Pop())
	}

	fmt.Printf("Reversed string: %s: ", revString)
	return revString
}
