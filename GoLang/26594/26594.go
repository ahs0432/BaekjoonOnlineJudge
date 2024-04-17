package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	now := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != now {
			break
		}
		count++
	}
	fmt.Println(count)
}
