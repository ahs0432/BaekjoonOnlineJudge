package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	s := "*"
	for i := 1; i < n; i++ {
		s += " *"
	}

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(" " + s)
		} else {
			fmt.Println(s)
		}
	}
}
