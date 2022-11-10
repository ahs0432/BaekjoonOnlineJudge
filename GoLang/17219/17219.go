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

	memo := map[string]string{}

	for i := 0; i < n; i++ {
		var site, passwd string
		fmt.Fscanln(reader, &site, &passwd)
		memo[site] = passwd
	}

	for i := 0; i < m; i++ {
		var site string
		fmt.Fscanln(reader, &site)
		fmt.Fprintln(writer, memo[site])
	}
	writer.Flush()
}
