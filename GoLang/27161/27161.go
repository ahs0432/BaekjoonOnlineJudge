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

	now := 1
	add := 1
	ans := "NO"
	var s string
	var x int
	for i := 0; i < n; i++ {
		ans = "NO"
		fmt.Fscanln(reader, &s, &x)

		if s == "HOURGLASS" && now == x {
			ans = "NO"
		} else if s == "HOURGLASS" {
			add *= -1
		} else {
			if now == x {
				ans = "YES"
			}
		}
		fmt.Fprintln(writer, now, ans)
		now += add

		if now == 13 {
			now = 1
		} else if now == 0 {
			now = 12
		}
	}
	writer.Flush()
}
