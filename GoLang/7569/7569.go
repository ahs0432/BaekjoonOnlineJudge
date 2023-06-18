package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, h int
	fmt.Fscanln(reader, &n, &m, &h)

	box := make([][][]int, h)
	boxtmp := make([][][]int, h)

	for i := 0; i < h; i++ {
		box[i] = make([][]int, m)
		boxtmp[i] = make([][]int, m)

		for j := 0; j < m; j++ {
			box[i][j] = make([]int, n)
			boxtmp[i][j] = make([]int, n)

			for k := 0; k < n; k++ {
				if k != n-1 {
					fmt.Fscan(reader, &box[i][j][k])
				} else {
					fmt.Fscanln(reader, &box[i][j][k])
				}
			}
			copy(boxtmp[i][j], box[i][j])
		}
	}

	moves := [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	count := 0

	queue := make([][]int, 0)

	for i := 0; i < h; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if box[i][j][k] == 1 {
					queue = append(queue, []int{i, j, k, 0})
				}
			}
		}
	}

	for len(queue) != 0 {
		x := queue[0][0]
		y := queue[0][1]
		z := queue[0][2]
		tmpCount := queue[0][3] + 1

		queue = queue[1:]
		for _, move := range moves {
			moveX := x + move[0]
			moveY := y + move[1]
			moveZ := z + move[2]

			if moveX < 0 || moveX == h || moveY < 0 || moveY == m || moveZ < 0 || moveZ == n {
				continue
			}

			if box[moveX][moveY][moveZ] == 0 {
				queue = append(queue, []int{moveX, moveY, moveZ, tmpCount})
				box[moveX][moveY][moveZ] = 1

				if count < tmpCount {
					count = tmpCount
				}
			}

		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if box[i][j][k] == 0 {
					fmt.Println(-1)
					return
				}
			}
		}
	}

	fmt.Println(count)
}
