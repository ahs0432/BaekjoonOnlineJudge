package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, b int
	fmt.Fscanln(reader, &r, &b)

	var l, w int
	for i := 3; ; i++ {
		for j := i; j >= 3; j-- {
			l = (i - 2) * (j - 2)
			w = i*j - l

			if w == r && l == b {
				fmt.Println(i, j)
				return
			}
		}
	}
}
