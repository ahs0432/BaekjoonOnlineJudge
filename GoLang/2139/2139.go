package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	dates := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var d, m, y int
	var ans int
	for {
		fmt.Fscanln(reader, &d, &m, &y)

		if d == 0 && m == 0 && y == 0 {
			break
		}

		ans = 0

		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			for i := 0; i < m; i++ {
				ans += dates[i]
			}
			ans += d

			if m > 2 {
				ans++
			}
		} else {
			for i := 0; i < m; i++ {
				ans += dates[i]
			}
			ans += d
		}

		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
