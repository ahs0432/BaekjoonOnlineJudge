package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	p := make([]int, 6)
	p[0] = 1
	for i := 1; i <= 5; i++ {
		p[i] = i * 1 * p[i-1]
	}

	var n int
	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		sum := 0
		for i := 1; n != 0; i++ {
			sum += (n % 10) * p[i]
			n /= 10
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
