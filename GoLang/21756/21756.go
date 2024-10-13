package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]int, n)
	for i := 1; i <= n; i++ {
		table[i-1] = i
	}

	tmp := []int{}
	for len(table) != 1 {
		for i, num := range table {
			if i%2 == 1 {
				tmp = append(tmp, num)
			}
		}

		table = tmp
		tmp = []int{}
	}

	fmt.Println(table[0])
}
