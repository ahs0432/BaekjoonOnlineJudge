package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, sum int
	var table []int
	for {
		fmt.Fscanln(reader, &n)
		if n == -1 {
			break
		}

		sum = 0
		table = make([]int, 0)
		for i := 2; i < n/2+1; i++ {
			if n%i == 0 {
				sum += i
				table = append(table, i)
			}
		}

		if sum+1 == n {
			fmt.Fprintf(writer, "%d = 1", n)
			for _, j := range table {
				fmt.Fprintf(writer, " + %d", j)
			}
			fmt.Fprintln(writer)
		} else {
			fmt.Fprintf(writer, "%d is NOT perfect.\n", n)
		}
	}
	writer.Flush()
}
