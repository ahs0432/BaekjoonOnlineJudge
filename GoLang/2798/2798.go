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

	table := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	max := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				tmp := table[i] + table[j] + table[k]

				if tmp > m {
					continue
				} else if tmp == m {
					fmt.Println(tmp)
					return
				}

				if max < tmp {
					max = tmp
				}
			}
		}
	}

	fmt.Println(max)
}
