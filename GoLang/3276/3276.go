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

	a, b := 1, 1
	for a*b < n {
		if a > b {
			b++
		} else {
			a++
		}
	}
	fmt.Println(a, b)
}
