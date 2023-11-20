package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		for len(n) >= 2 {
			sum = 0
			for _, c := range n {
				sum += int(c - 48)
			}
			n = strconv.Itoa(sum)
		}

		fmt.Fprintln(writer, n)
	}
	writer.Flush()
}
