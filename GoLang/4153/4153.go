package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var x, y, z int
		fmt.Fscanln(reader, &x, &y, &z)

		if x == 0 && y == 0 && z == 0 {
			break
		}

		if x > y {
			tmp := x
			x = y
			y = tmp
		}

		if y > z {
			tmp := y
			y = z
			z = tmp
		}

		if z*z == (x*x)+(y*y) {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
	}
}
