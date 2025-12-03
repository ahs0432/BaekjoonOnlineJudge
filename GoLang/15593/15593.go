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

	table := make([]int, 1000)
	works := make([][]int, n)

	for i := 0; i < n; i++ {
		works[i] = make([]int, 2)
		fmt.Fscanln(reader, &works[i][0], &works[i][1])

		for j := works[i][0]; j < works[i][1]; j++ {
			table[j]++
		}
	}

	ans := 0
	var time int
	for _, work := range works {
		for i := work[0]; i < work[1]; i++ {
			table[i]--
		}
		time = 0
		for i := 0; i < len(table); i++ {
			if table[i] > 0 {
				time++
			}
		}

		ans = max(time, ans)
		for i := work[0]; i < work[1]; i++ {
			table[i]++
		}
	}
	fmt.Println(ans)
}
