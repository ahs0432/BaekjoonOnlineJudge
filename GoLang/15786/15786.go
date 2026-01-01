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

	var str string
	fmt.Fscanln(reader, &str)

	var count int
	var tmp string
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &tmp)

		count = 0
		for _, c := range tmp {
			if byte(c) == str[count] {
				count++
				if n <= count {
					break
				}
			}
		}

		if n == count {
			fmt.Fprintln(writer, "true")
		} else {
			fmt.Fprintln(writer, "false")
		}
	}
	writer.Flush()
}
