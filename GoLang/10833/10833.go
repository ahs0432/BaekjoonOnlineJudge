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

	var a, b int
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a < b {
			sum += b % a
		} else if a > b {
			sum += b
		}
	}

	fmt.Println(sum)
}
