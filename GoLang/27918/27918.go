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

	x, y := 0, 0

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		if !(x >= y+2) && !(x+2 <= y) {
			if tmp == "D" {
				x++
			} else {
				y++
			}
		}
	}
	fmt.Printf("%d:%d\n", x, y)
}
