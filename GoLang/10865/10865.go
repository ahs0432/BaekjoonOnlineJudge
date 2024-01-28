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
	fmt.Fscanln(reader, &n, &m)

	t := make([]int, n)

	var tmp1, tmp2 int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &tmp1, &tmp2)
		t[tmp1-1]++
		t[tmp2-1]++
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, t[i])
	}

	writer.Flush()
}
