package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a1, b1, c1 float64
	fmt.Fscanln(reader, &a1, &b1, &c1)

	a2 := (a1 + b1 - c1) / 2
	b2 := a1 - a2
	c2 := b1 - a2

	if a2 > 0 && b2 > 0 && c2 > 0 {
		fmt.Println(1)
		fmt.Printf("%.1f %.1f %.1f\n", a2, b2, c2)
	} else {
		fmt.Println(-1)
	}
}
