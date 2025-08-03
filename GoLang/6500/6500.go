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

	var n int
	var table []int
	var s string
	var check bool
	for {
		fmt.Fscanln(reader, &s)
		n, _ = strconv.Atoi(s)

		if n == 0 {
			break
		}

		table = []int{}
		table = append(table, n)
		n *= n

		for {
			s = strconv.Itoa(n)
			for len(s) < 8 {
				s = "0" + s
			}
			n, _ = strconv.Atoi(s[2:6])

			check = false
			for _, item := range table {
				if item == n {
					check = true
					break
				}
			}

			if check {
				break
			}

			table = append(table, n)
			n *= n
		}
		fmt.Fprintln(writer, len(table))
	}
	writer.Flush()
}
