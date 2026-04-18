package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	iTable := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &iTable[i])
	}

	jTable := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &jTable[i])
	}

	var tmp int
	var tmpTable []int
	for i := 0; i < n; i++ {
		tmpTable = make([]int, m)
		tmp = iTable[i]

		for j := 0; j < m; j++ {
			if tmp == jTable[j] {
				tmp = 0
			} else {
				tmp = 1
			}
			tmpTable[j] = tmp
		}
		jTable = tmpTable
	}

	if len(jTable) > 0 {
		fmt.Fprintln(writer, jTable[len(jTable)-1])
	}
	writer.Flush()
}
