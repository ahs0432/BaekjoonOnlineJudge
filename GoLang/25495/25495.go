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

	p := 0
	b := 2
	ans := 0
	var a int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)

		if p == a {
			b *= 2
		} else {
			b = 2
		}

		p = a
		ans += b

		if ans >= 100 {
			ans = 0
			p = 0
			b = 2
		}
	}
	fmt.Println(ans)
}
