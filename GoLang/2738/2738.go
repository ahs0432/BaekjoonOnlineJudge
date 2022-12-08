package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y int
	fmt.Fscanln(reader, &x, &y)

	answer := make([][]int, x)

	for i := 0; i < x; i++ {
		answer[i] = make([]int, y)
		for j := 0; j < y; j++ {
			if j == y-1 {
				fmt.Fscanln(reader, &answer[i][j])
			} else {
				fmt.Fscan(reader, &answer[i][j])
			}
		}
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			var tmp int
			if j == y-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			answer[i][j] += tmp
		}
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(answer[i][j], " ")
		}
		fmt.Println()
	}
}
