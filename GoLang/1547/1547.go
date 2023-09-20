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

	now := 1
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a == now {
			now = b
		} else if b == now {
			now = a
		}
	}

	fmt.Println(now)
}
