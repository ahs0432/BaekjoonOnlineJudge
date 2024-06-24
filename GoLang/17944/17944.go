package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscanln(reader, &n, &t)

	x := 1
	count := 0
	for i := 0; i < t; i++ {
		count += x
		if count == 2*n {
			x = -1
		}
		if count == 1 {
			x = 1
		}
	}

	fmt.Println(count)
}
