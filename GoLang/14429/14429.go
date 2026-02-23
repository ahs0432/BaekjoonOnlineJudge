package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var j, m, next, ans int
	now := 2147483647
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &j, &m)
		next = ((j-1)/(m+1) + 1) * 2
		if next < now {
			now = next
			ans = i
		}
	}
	fmt.Println(ans, now)
}
