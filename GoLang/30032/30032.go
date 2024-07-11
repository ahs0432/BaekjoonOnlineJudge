package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, d int
	fmt.Fscanln(reader, &n, &d)

	var tmp string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		for j := 0; j < len(tmp); j++ {
			if d == 1 {
				if tmp[j] == 'd' {
					fmt.Fprint(writer, "q")
				} else if tmp[j] == 'b' {
					fmt.Fprint(writer, "p")
				} else if tmp[j] == 'q' {
					fmt.Fprint(writer, "d")
				} else if tmp[j] == 'p' {
					fmt.Fprint(writer, "b")
				}
			} else {
				if tmp[j] == 'd' {
					fmt.Fprint(writer, "b")
				} else if tmp[j] == 'b' {
					fmt.Fprint(writer, "d")
				} else if tmp[j] == 'q' {
					fmt.Fprint(writer, "p")
				} else if tmp[j] == 'p' {
					fmt.Fprint(writer, "q")
				}
			}
		}

		fmt.Fprintln(writer)
	}

	writer.Flush()
}
