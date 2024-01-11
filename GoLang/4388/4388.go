package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	count := 0
	for {
		fmt.Fscanln(reader, &a, &b)

		if a == 0 && b == 0 {
			break
		}

		count = 0
		c = 0
		for a != 0 || b != 0 {
			if a%10+b%10+c >= 10 {
				count++
				c = 1
			} else {
				c = 0
			}
			a /= 10
			b /= 10
		}

		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
