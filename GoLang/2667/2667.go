package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]bool, n)
	check := make([][]bool, n)

	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
		check[i] = make([]bool, n)

		var tmp string
		fmt.Fscanln(reader, &tmp)

		for j, t := range tmp {
			if t == '1' {
				table[i][j] = true
			}
		}
	}

	answer := make([]int, 0)
	connCheck := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i][j] && !check[i][j] {
				now := 1
				q := [][]int{{i, j}}
				check[i][j] = true

				for len(q) != 0 {
					x := q[0][0]
					y := q[0][1]

					q = q[1:]

					for _, c := range connCheck {
						nowX := x + c[0]
						nowY := y + c[1]

						if nowX < 0 || nowX == n || nowY < 0 || nowY == n || !table[nowX][nowY] || check[nowX][nowY] {
							continue
						}

						now++
						check[nowX][nowY] = true
						q = append(q, []int{nowX, nowY})
					}
				}
				answer = append(answer, now)
			}
		}
	}

	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})

	fmt.Fprintln(writer, len(answer))
	for _, ans := range answer {
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
