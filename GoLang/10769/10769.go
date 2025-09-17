package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	happy := 0
	sad := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] == ':' && s[i+1] == '-' {
			switch s[i+2] {
			case ')':
				happy++
			case '(':
				sad++
			}
		}
	}

	if happy == 0 && sad == 0 {
		fmt.Println("none")
	} else if happy == sad {
		fmt.Println("unsure")
	} else if happy > sad {
		fmt.Println("happy")
	} else {
		fmt.Println("sad")
	}
}
