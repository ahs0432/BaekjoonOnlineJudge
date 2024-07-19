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

	var a, b, c, d int
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &c, &d)

		if i > 0 {
			if a == c && a != 0 {
				cnt++
			}
			if b == d && b != 0 {
				cnt++
			}
		}

		if c == d && c != 0 {
			cnt++
		}
		a = c
		b = d
	}
	fmt.Println(cnt)
}
