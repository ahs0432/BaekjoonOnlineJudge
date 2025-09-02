package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &table[i])
	}

	var tmp int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if table[j] > table[j+1] {
				tmp = table[j+1]
				table[j+1] = table[j]
				table[j] = tmp
				for k := 0; k < 5; k++ {
					fmt.Fprint(writer, table[k], " ")
				}
				fmt.Fprintln(writer)
			}
		}
	}
	writer.Flush()
}
