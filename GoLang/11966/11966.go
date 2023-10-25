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

	s := 1
	for i := 0; i < 31; i++ {
		if s == n {
			fmt.Println(1)
			return
		} else if s > n {
			break
		}

		s *= 2
	}
	fmt.Println(0)
}
