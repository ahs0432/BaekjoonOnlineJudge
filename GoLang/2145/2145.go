package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, tmp int

	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		for n > 9 {
			tmp = 0
			for n != 0 {
				tmp += n % 10
				n /= 10
			}
			n = tmp
		}

		fmt.Fprintln(writer, n)
	}
	writer.Flush()
}
