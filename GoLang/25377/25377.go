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

	r := -1

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		if a <= b && (b < r || r == -1) {
			r = b
		}
	}

	fmt.Println(r)
}
