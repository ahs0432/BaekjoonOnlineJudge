package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, z int
	for {
		fmt.Fscanln(reader, &x, &y, &z)

		if x == 0 && y == 0 && z == 0 {
			break
		} else if x == y && y == z {
			fmt.Println("Equilateral")
		} else if x >= y+z || y >= x+z || z >= x+y {
			fmt.Println("Invalid")
		} else if x == y || x == z || y == z {
			fmt.Println("Isosceles")
		} else if x != y && x != z && y != z {
			fmt.Println("Scalene")
		}
	}
}
