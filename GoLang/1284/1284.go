package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n string
	var sum int
	for {
		fmt.Fscanln(reader, &n)

		if n == "0" {
			break
		}

		sum = 1
		for _, c := range n {
			if c == 48 {
				sum += 5
			} else if c == 49 {
				sum += 3
			} else {
				sum += 4
			}
		}

		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
