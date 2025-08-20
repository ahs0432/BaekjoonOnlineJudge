package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	max := 0
	s := make([]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &s[i])

		if len(s[i]) > max {
			max = len(s[i])
		}
	}

	for i := 0; i < max; i++ {
		for j := 0; j < 5; j++ {
			if len(s[j]) > i {
				fmt.Fprint(writer, string(s[j][i]))
			}
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
