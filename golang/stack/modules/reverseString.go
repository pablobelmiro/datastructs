package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RevString(s *Stack) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		s.El = s.El[:0]
		fmt.Println("Enter a Word(max:10)(stop to break):")
		var revString string

		if !reader.Scan() {
			break
		}

		word := strings.TrimSpace(reader.Text())

		_, err := strconv.Atoi(word)
		if err == nil {
			fmt.Println("Only string is accepted!")
			continue
		}

		if word == "stop" {
			break
		}

		for _, c := range word {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				s.Push(int(c))
			} else {
				fmt.Println("Only alphabetic strings are accepted!")
			}
		}

		for s.Len() > 0 {
			revString += string(s.Pop())
		}

		fmt.Printf("Reversed string: %s\n", revString)
	}
}
