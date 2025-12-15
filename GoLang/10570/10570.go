package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, tmp, ans int
	fmt.Fscanln(reader, &t)

	var table []int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		table = make([]int, 1001)

		ans = 0
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &tmp)
			table[tmp]++

			if table[ans] < table[tmp] {
				ans = tmp
			} else if table[ans] == table[tmp] && ans > tmp {
				ans = tmp
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
