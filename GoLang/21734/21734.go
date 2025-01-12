package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	fmt.Fscanln(reader, &s)

	var sum, tmp int
	for _, c := range s {
		sum = 0
		tmp = int(c)

		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}

		for i := 0; i < sum; i++ {
			fmt.Fprint(writer, string(c))
		}

		fmt.Fprintln(writer)
	}

	writer.Flush()
}
