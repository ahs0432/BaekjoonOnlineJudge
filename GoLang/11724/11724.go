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

	table := make([][]int, n)
	check := make([]bool, n)

	for i := 0; i < n; i++ {
		table[i] = []int{}
	}

	for i := 0; i < m; i++ {
		var tmp1, tmp2 int
		fmt.Fscanln(reader, &tmp1, &tmp2)

		tmp1 -= 1
		tmp2 -= 1

		table[tmp1] = append(table[tmp1], tmp2)
		table[tmp2] = append(table[tmp2], tmp1)
	}

	count := 0
	for i := 0; i < n; i++ {
		if check[i] {
			continue
		}
		count++

		queue := []int{i}
		check[i] = true

		for len(queue) != 0 {
			for _, j := range table[queue[0]] {
				if check[j] {
					continue
				}

				queue = append(queue, j)
				check[j] = true
			}
			queue = queue[1:]
		}
	}

	fmt.Println(count)
}
