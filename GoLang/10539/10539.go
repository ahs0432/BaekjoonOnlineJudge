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

	var tmp int
	sum, res := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		res = tmp*(i+1) - sum
		sum += res
		fmt.Fprint(writer, res, " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
