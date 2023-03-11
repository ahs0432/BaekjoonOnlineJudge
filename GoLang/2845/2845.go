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

	sum := n * m

	for i := 0; i < 5; i++ {
		var tmp int
		if i != 4 {
			fmt.Fscan(reader, &tmp)
		} else {
			fmt.Fscanln(reader, &tmp)
		}
		fmt.Fprint(writer, tmp-sum, " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
