package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s, sum int
	for {
		fmt.Fscanln(reader, &s)

		if s == 0 {
			break
		}

		fmt.Fprint(writer, s, " ")
		for s >= 10 {
			sum = 1
			for s != 0 {
				sum *= (s % 10)
				s /= 10
			}
			fmt.Fprint(writer, sum, " ")
			s = sum
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
