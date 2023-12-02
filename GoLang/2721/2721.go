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

	var tmp, sum, ans int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		sum, ans = 0, 0
		for j := 1; j <= tmp; j++ {
			for k := 1; k <= j+1; k++ {
				sum += k
			}
			ans += j * sum
			sum = 0
		}

		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
