package main

import (
	"bufio"
	"fmt"
	"os"
)

func createBox() [][]int {
	box := make([][]int, 100)
	for i := 0; i < 100; i++ {
		box[i] = make([]int, 100)
	}
	return box
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var m, n int
	var count int
	var box [][]int
	for i := 0; i < t; i++ {
		count = 0
		box = createBox()

		fmt.Fscanln(reader, &m, &n)

		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if k == n-1 {
					fmt.Fscanln(reader, &box[j][k])

				} else {
					fmt.Fscan(reader, &box[j][k])
				}
			}
		}

		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				if box[k][j] == 1 {
					for l := k + 1; l < m; l++ {
						if box[l][j] == 0 {
							count++
						}
					}
				}
			}
		}
		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
