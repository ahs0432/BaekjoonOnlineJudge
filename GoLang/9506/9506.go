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
	var tmp []int
	for {
		fmt.Fscanln(reader, &n)
		if n == -1 {
			break
		}

		sum = 0
		tmp = []int{}
		for i := 1; i < n; i++ {
			if n%i == 0 {
				tmp = append(tmp, i)
				sum += i
			}
		}

		if n == sum {
			fmt.Fprintf(writer, "%d = %d", n, 1)
			for i := 1; i < len(tmp); i++ {
				fmt.Fprintf(writer, " + %d", tmp[i])
			}
			fmt.Fprintln(writer)
		} else {
			fmt.Fprintf(writer, "%d is NOT perfect.\n", n)
		}
	}
	writer.Flush()
}
