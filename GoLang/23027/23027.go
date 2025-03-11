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
	fmt.Fscanln(reader, &s)

	if strings.Contains(s, "A") {
		s = strings.ReplaceAll(s, "B", "A")
		s = strings.ReplaceAll(s, "C", "A")
		s = strings.ReplaceAll(s, "D", "A")
		s = strings.ReplaceAll(s, "F", "A")
	} else if strings.Contains(s, "B") {
		s = strings.ReplaceAll(s, "C", "B")
		s = strings.ReplaceAll(s, "D", "B")
		s = strings.ReplaceAll(s, "F", "B")
	} else if strings.Contains(s, "C") {
		s = strings.ReplaceAll(s, "D", "C")
		s = strings.ReplaceAll(s, "F", "C")
	} else {
		l := len(s)
		s = ""
		for i := 0; i < l; i++ {
			s += "A"
		}
	}
	fmt.Println(s)
}
