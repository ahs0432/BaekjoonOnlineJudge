package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	table := make([]bool, n)
	ans := n
	var op, x int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &op)

		if op < 3 {
			fmt.Fscan(reader, &x)
			x--
		}

		if op == 1 && !table[x] {
			table[x] = true
			ans--
		}

		if op == 2 && table[x] {
			table[x] = false
			ans++
		}

		if op == 3 {
			fmt.Fprintln(writer, ans)
		}
	}
	writer.Flush()
}
