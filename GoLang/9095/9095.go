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

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)

		table := make([]int, tmp+1)

		if tmp < 3 {
			fmt.Println(tmp)
			continue
		}

		table[1] = 1
		table[2] = 2
		table[3] = 4

		for j := 4; j <= tmp; j++ {
			table[j] = table[j-1] + table[j-2] + table[j-3]
		}

		fmt.Println(table[tmp])
	}
}
