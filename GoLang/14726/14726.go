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

	var sum int
	var tmp byte
	var id []byte
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &id)

		for i := 0; i < 16; i += 2 {
			tmp = (id[i] - '0') * 2
			if tmp > 0 {
				id[i] = (tmp/10 + tmp%10) + '0'
			}
		}

		sum = 0
		for _, c := range id {
			sum += int(c - '0')
		}

		if sum%10 == 0 {
			fmt.Fprintln(writer, "T")
		} else {
			fmt.Fprintln(writer, "F")
		}
	}
	writer.Flush()
}
