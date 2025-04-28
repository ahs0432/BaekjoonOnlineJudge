package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([][]bool, 1002)
	check := make([][]bool, 1002)

	for i := 0; i < 1002; i++ {
		table[i] = make([]bool, 1002)
		check[i] = make([]bool, 1002)
	}

	var x, y, n int
	fmt.Fscanln(reader, &x, &y, &n)

	x += 500
	y += 500

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		table[b+500][a+500] = true
	}

	stack := [][]int{{0, 500, 500}}

	var nowX, nowY, nowSize int
	var nextX, nextY, nextSize int

	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}

	for len(stack) > 0 {
		nowSize = stack[0][0]
		nowY = stack[0][1]
		nowX = stack[0][2]
		stack = stack[1:len(stack)]

		if check[nowY][nowX] {
			continue
		}

		check[nowY][nowX] = true

		nextSize = nowSize + 1
		for i := 0; i < 4; i++ {
			nextY = nowY + dy[i]
			nextX = nowX + dx[i]

			if nextY > 1000 || nextY < 0 || nextX > 1000 || nextX < 0 || table[nextY][nextX] || check[nextY][nextX] {
				continue
			}

			if nextY == y && nextX == x {
				fmt.Println(nextSize)
				return
			}
			stack = append(stack, []int{nextSize, nextY, nextX})
		}
	}
}
