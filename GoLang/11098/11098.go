package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, p int
	fmt.Fscanln(reader, &n)

	var c, max int
	var name, ans string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &p)
		c = 0
		max = 0

		for j := 0; j < p; j++ {
			fmt.Fscanln(reader, &c, &name)

			if max < c {
				ans = name
				max = c
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
