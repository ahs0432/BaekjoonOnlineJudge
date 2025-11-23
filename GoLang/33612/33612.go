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

	y := 2024
	m := 1 + (n * 7)

	for m > 12 {
		y++
		m -= 12
	}

	fmt.Println(y, m)
}
