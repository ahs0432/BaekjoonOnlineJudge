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

	maze := make([][]bool, x)

	for i := 0; i < x; i++ {
		maze[i] = make([]bool, y)
		var tmp string
		fmt.Fscanln(reader, &tmp)
		for j := 0; j < y; j++ {
			if tmp[j] == '1' {
				maze[i][j] = true
			}
		}
	}

	if x == 1 && y == 1 {
		fmt.Println(1)
		return
	}

	queue := [][]int{{0, 0, 1}}
	maze[0][0] = false
	moveQ := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) != 0 {
		nowX := queue[0][0]
		nowY := queue[0][1]
		count := queue[0][2] + 1
		queue = queue[1:]

		for _, move := range moveQ {
			moveX := nowX + move[0]
			moveY := nowY + move[1]

			if moveX < 0 || moveY < 0 || moveX > x-1 || moveY > y-1 || !maze[moveX][moveY] {
				continue
			}

			if moveX == x-1 && moveY == y-1 {
				fmt.Println(count)
				return
			}

			queue = append(queue, []int{moveX, moveY, count})
			maze[moveX][moveY] = false
		}
	}
}
