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
	fmt.Fscanln(reader, &n)

	var tmp string
	var chars map[byte]bool
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		if len(tmp) != 7 {
			fmt.Fprintln(writer, 0)
			continue
		}

		chars = map[byte]bool{}
		for j := 0; j < len(tmp); j++ {
			chars[tmp[j]] = true
		}

		if len(chars) != 2 {
			fmt.Fprintln(writer, 0)
			continue
		}

		if tmp[0] == tmp[1] && tmp[1] == tmp[4] && tmp[2] == tmp[3] && tmp[3] == tmp[5] && tmp[5] == tmp[6] {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}

	writer.Flush()
}
