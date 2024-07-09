package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(num, base int) int {
	ret := 0

	for num != 0 {
		ret += num % base
		num /= base
	}

	return ret
}

func main() {
	writer := bufio.NewWriter(os.Stdout)

	for i := 1001; i < 10000; i++ {
		tmp := check(i, 10)
		if tmp == check(i, 12) && tmp == check(i, 16) {
			fmt.Fprintln(writer, i)
		}
	}
	writer.Flush()
}
