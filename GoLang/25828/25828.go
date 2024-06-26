package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var g, p, t int
	fmt.Fscanln(reader, &g, &p, &t)

	a, b := g*p, g+p*t
	if a < b {
		fmt.Println(1)
	} else if a > b {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
