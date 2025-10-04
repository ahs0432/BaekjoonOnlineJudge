package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(table []int, n int) int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			} else {
				if table[i]*2 == table[j] {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var table []int = []int{0}
	for table[0] != -1 {
		table = make([]int, 16)
		for i := 0; i < 16; i++ {
			fmt.Fscan(reader, &table[i])
			if table[0] == -1 {
				break
			}

			if table[i] == 0 {
				fmt.Fprintln(writer, check(table, i))
				break
			}
		}

	}

	writer.Flush()
}
