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

	changeA := make([][]int, n)
	changeB := make([][]int, n)

	for i := 0; i < n; i++ {
		changeA[i] = make([]int, m)
		changeB[i] = make([]int, m)

		var tmp string
		fmt.Fscanln(reader, &tmp)

		for j, t := range tmp {
			if (i+j)%2 == 0 {
				if string(t) == "W" {
					changeA[i][j] = 1
				} else {
					changeB[i][j] = 1
				}
			} else {
				if string(t) == "B" {
					changeA[i][j] = 1
				} else {
					changeB[i][j] = 1
				}
			}
		}
	}

	min := 8 * 8

	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			tmpA := 0
			tmpB := 0

			for p := i; p < i+8; p++ {
				for q := j; q < j+8; q++ {
					tmpA += changeA[p][q]
					tmpB += changeB[p][q]
				}
			}

			if min > tmpA {
				min = tmpA
			}

			if min > tmpB {
				min = tmpB
			}
		}
	}

	fmt.Println(min)
}
