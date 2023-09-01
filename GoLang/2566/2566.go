package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	max, tmp, x, y := 0, 0, 0, 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j == 8 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			if tmp > max {
				max = tmp
				x = i
				y = j
			}
		}
	}
	fmt.Println(max)
	fmt.Println(x+1, y+1)
}
