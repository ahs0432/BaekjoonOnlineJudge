package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, c, n int
	fmt.Fscanln(reader, &r, &c, &n)

	x := r / n
	if r%n != 0 {
		x++
	}

	y := c / n
	if c%n != 0 {
		y++
	}

	fmt.Println(x * y)
}
