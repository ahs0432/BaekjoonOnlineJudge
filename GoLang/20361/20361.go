package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x, k int
	fmt.Fscanln(reader, &n, &x, &k)

	var a, b int
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &a, &b)
		if a == x {
			x = b
		} else if b == x {
			x = a
		}
	}
	fmt.Println(x)
}
