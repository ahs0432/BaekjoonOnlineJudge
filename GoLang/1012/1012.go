package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	var table [50][50]bool

	for i := 0; i < t; i++ {
		var m, n, k int
		fmt.Fscanln(reader, &m, &n, &k)

		for j := 0; j < k; j++ {
			var x, y int
			fmt.Fscanln(reader, &x, &y)
			table[x][y] = true
		}

		count := 0
		for x := 0; x < m; x++ {
			for y := 0; y < n; y++ {
				if table[x][y] {
					count++
					queue := [][]int{{x, y}}
					table[x][y] = false

					for len(queue) != 0 {
						qx := queue[0][0]
						qy := queue[0][1]
						queue = queue[1:]

						for j := 0; j < 4; j++ {
							if qx+dx[j] < 0 || qx+dx[j] == m || qy+dy[j] < 0 || qy+dy[j] == n || !table[qx+dx[j]][qy+dy[j]] {
								continue
							}

							queue = append(queue, []int{qx + dx[j], qy + dy[j]})
							table[qx+dx[j]][qy+dy[j]] = false
						}
					}
				}
			}
		}

		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
