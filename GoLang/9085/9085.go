package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)

		sum := 0
		var a int

		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &a)
			} else {
				fmt.Fscan(reader, &a)
			}
			sum += a
		}

		fmt.Println(sum)
	}
}
