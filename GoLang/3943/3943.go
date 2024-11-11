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

	var tmp, max int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		max = tmp

		for tmp != 1 {
			if tmp%2 == 0 {
				tmp = tmp / 2
			} else {
				tmp = tmp*3 + 1
			}

			if max < tmp {
				max = tmp
			}
		}
		fmt.Fprintln(writer, max)
	}
	writer.Flush()
}
