package main

import (
	"bufio"
	"fmt"
	"os"
)

func checking(table [][]int, startX int, endX int, startY int, endY int) []int {
	count := []int{0, 0, 0}
	check, now := true, table[startX][startY]
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			if now != table[i][j] {
				check = false
				break
			}
		}

		if !check {
			break
		}
	}

	if check {
		count[now+1] += 1
		return count
	}

	div := (endX - startX) / 3
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			tmp := checking(table, startX+(div*(i-1)), startX+(div*i), startY+(div*(j-1)), startY+(div*j))
			count[0] += tmp[0]
			count[1] += tmp[1]
			count[2] += tmp[2]
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]int, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &table[i][j])
		}
	}

	for _, tmp := range checking(table, 0, n, 0, n) {
		fmt.Println(tmp)
	}
}
