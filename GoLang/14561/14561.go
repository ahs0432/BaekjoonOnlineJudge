package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b int
	var table []int
	var check bool
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		table = []int{}
		for a != 0 {
			table = append(table, a%b)
			a /= b
		}

		check = true
		for k, l := 0, len(table)-1; k <= l; {
			if table[k] == table[l] {
				k++
				l--
			} else {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
	writer.Flush()
}
