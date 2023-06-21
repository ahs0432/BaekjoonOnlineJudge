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

	box := make([][]int, m)
	boxtmp := make([][]int, m)

	for i := 0; i < m; i++ {
		box[i] = make([]int, n)
		boxtmp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if j != n-1 {
				fmt.Fscan(reader, &box[i][j])
			} else {
				fmt.Fscanln(reader, &box[i][j])
			}

		}
		copy(boxtmp[i], box[i])
	}

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0

	queue := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if box[i][j] == 1 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}

	for len(queue) != 0 {
		x := queue[0][0]
		y := queue[0][1]
		tmpCount := queue[0][2] + 1

		queue = queue[1:]
		for _, move := range moves {
			moveX := x + move[0]
			moveY := y + move[1]

			if moveX < 0 || moveX == m || moveY < 0 || moveY == n {
				continue
			}

			if box[moveX][moveY] == 0 {
				queue = append(queue, []int{moveX, moveY, tmpCount})
				box[moveX][moveY] = 1

				if count < tmpCount {
					count = tmpCount
				}
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if box[i][j] == 0 {
				fmt.Println(-1)
				return
			}
		}
	}

	fmt.Println(count)
}
