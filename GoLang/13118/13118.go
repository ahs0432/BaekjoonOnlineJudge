package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	p := make([]int, 4)
	fmt.Fscanln(reader, &p[0], &p[1], &p[2], &p[3])

	var x, y, z int
	fmt.Fscanln(reader, &x, &y, &z)

	ans := -1

	for i, n := range p {
		if n == x {
			ans = i
		}
	}
	fmt.Println(ans + 1)
}
