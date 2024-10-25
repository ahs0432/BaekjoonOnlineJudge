package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, m int
	fmt.Fscanln(reader, &a, &b, &c, &m)

	now := 0
	ans := 0
	for i := 0; i < 24; i++ {
		if now+a <= m {
			now += a
			ans += b
		} else {
			if now-c >= 0 {
				now -= c
			} else {
				now = 0
			}
		}
	}
	fmt.Println(ans)
}
