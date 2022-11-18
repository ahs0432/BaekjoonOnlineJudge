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

	connection := map[int][]int{}

	friends := make([][]int, n)

	for i := 0; i < n; i++ {
		friends[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				friends[i][j] = 2147483647
			}
		}
	}

	for i := 0; i < m; i++ {
		var tmpi, tmpj int
		fmt.Fscanln(reader, &tmpi, &tmpj)
		tmpi--
		tmpj--

		connection[tmpi] = append(connection[tmpi], tmpj)
		connection[tmpj] = append(connection[tmpj], tmpi)
	}

	minC := 0
	min := 2147483647
	for i := 0; i < n; i++ {
		total := 0
		queue := []int{i}

		for len(queue) != 0 {
			tmp := queue[0]
			queue = queue[1:]

			for _, q := range connection[tmp] {
				if friends[i][q] > (friends[i][tmp] + 1) {
					total += (friends[i][tmp] + 1)
					friends[i][q] = (friends[i][tmp] + 1)
					queue = append(queue, q)
				}
			}
		}

		if total < min {
			minC = i + 1
			min = total
		}
	}

	fmt.Println(minC)
}
