package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	r1 := b*2 + c*4
	r2 := a*2 + c*2
	r3 := a*4 + b*2

	fmt.Println(min(min(r1, r2), r3))
}
