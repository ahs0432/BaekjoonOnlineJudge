package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		var docs, want int
		fmt.Fscanln(reader, &docs, &want)

		table := make([]int, docs)
		for j := 0; j < docs-1; j++ {
			fmt.Fscan(reader, &table[j])
		}
		fmt.Fscanln(reader, &table[docs-1])

		for j := 0; j < docs; j++ {
			for k := j + 1; k < docs; k++ {
				if table[j] < table[k] {
					if j == want {
						want = docs - 1
					} else if j < want {
						want--
					}

					table = append(table, table[j])
					table = table[1:]
					j -= 1
					break
				}
			}
		}

		fmt.Println(want + 1)
	}
}
