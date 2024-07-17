package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, ans, tmp int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		ans = 0
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			ans += 100 * tmp
		}

		if ans == 0 {
			fmt.Fprintln(writer, "Equilibrium")
		} else if ans > 0 {
			fmt.Fprintln(writer, "Right")
		} else {
			fmt.Fprintln(writer, "Left")
		}
	}

	writer.Flush()
}
