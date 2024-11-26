package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	var tmp int
	for i := 0; i < t; i++ {
		tmp = 10
		fmt.Fscanln(reader, &n)

		for n > tmp {
			tmp *= 10
		}

		if n*n%tmp == n {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
