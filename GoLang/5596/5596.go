package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]int, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			var tmp int

			if j == 3 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			table[i] += tmp
		}
	}

	if table[0] < table[1] {
		fmt.Println(table[1])
	} else {
		fmt.Println(table[0])
	}
}
