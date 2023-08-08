package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w1, w2, h1, h2 int
	fmt.Fscanln(reader, &w1)
	fmt.Fscanln(reader, &h1)
	fmt.Fscanln(reader, &w2)
	fmt.Fscanln(reader, &h2)
	fmt.Println(4 + 2*max(w1, w2) + 2*(h1+h2))
}
