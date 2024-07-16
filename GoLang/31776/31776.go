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

	ans := 0

	table := make([]int, 3)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[0], &table[1], &table[2])

		if table[0]+table[1]+table[2] == -3 {
			continue
		}

		for j := 0; j < len(table); j++ {
			if table[j] == -1 {
				table[j] = 121
			}
		}

		if table[0] <= table[1] && table[1] <= table[2] {
			ans++
		}
	}
	fmt.Println(ans)
}
