package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	name := map[string]int{}
	number := map[int]string{}

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		name[tmp] = (i + 1)
		number[i+1] = tmp
	}

	for i := 0; i < m; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		if u, exist := name[tmp]; exist {
			fmt.Fprintln(writer, u)
		} else {
			num := 0

			for _, t := range tmp {
				num *= 10
				num += int(t) - 48
			}

			fmt.Fprintln(writer, number[num])
		}
	}

	writer.Flush()
}
