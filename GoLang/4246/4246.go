package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}

		var s string
		fmt.Fscanln(reader, &s)

		rows := len(s) / n
		arr := make([][]byte, rows)
		for i := range arr {
			arr[i] = make([]byte, n)
		}

		index := 0
		for i := 0; i < rows; i++ {
			if i%2 == 0 {
				for j := 0; j < n; j++ {
					arr[i][j] = s[index]
					index++
				}
			} else {
				for j := n - 1; j >= 0; j-- {
					arr[i][j] = s[index]
					index++
				}
			}
		}

		for j := 0; j < n; j++ {
			for i := 0; i < rows; i++ {
				fmt.Fprintf(writer, "%c", arr[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
