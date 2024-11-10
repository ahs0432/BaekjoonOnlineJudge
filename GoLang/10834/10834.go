package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscanln(reader, &m)

	var a, b, s int
	ans := 1
	now := 0
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b, &s)
		if s == 1 {
			now = 1 - now
		}
		ans = ans / a * b
	}
	fmt.Println(now, ans)
}
