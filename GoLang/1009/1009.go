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

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		tmpA := a % 10
		for j := 0; j < b-1; j++ {
			tmpA *= a
			if tmpA%10 == 0 {
				tmpA = 10
				break
			} else {
				tmpA %= 10
			}
		}

		if tmpA%10 == 0 {
			tmpA = 10
		}

		fmt.Fprintln(writer, tmpA)
	}

	writer.Flush()
}
