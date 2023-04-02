package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var c int
	fmt.Fscanln(reader, &c)

	c = (n + m) - (c * 2)

	if c < 0 {
		fmt.Println(n + m)
	} else {
		fmt.Println(c)
	}
}
