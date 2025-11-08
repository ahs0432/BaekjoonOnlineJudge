package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, p1, p2 int
	fmt.Fscanln(reader, &x, &y, &p1, &p2)

	for {
		if p1 > p2 {
			p2 += y
		} else if p1 < p2 {
			p1 += x
		} else {
			fmt.Println(p1)
			break
		}

		if p1 > 10000 || p2 > 10000 {
			fmt.Println(-1)
			break
		}
	}
}
