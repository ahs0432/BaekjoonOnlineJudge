package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, tmp int
	var t []int
	for {
		t = []int{}
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &tmp)

			if len(t) == 0 || t[len(t)-1] != tmp {
				t = append(t, tmp)
			}
		}

		for i := 0; i < len(t); i++ {
			fmt.Fprint(writer, t[i], " ")
		}
		fmt.Fprintln(writer, "$")
	}
	writer.Flush()
}
