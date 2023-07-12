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

	ans := 0
	for i := 0; i < n; i++ {
		var a, d, g int
		fmt.Fscanln(reader, &a, &d, &g)

		now := a * (d + g)
		if (d + g) == a {
			now *= 2
		}

		if ans < now {
			ans = now
		}
	}
	fmt.Println(ans)
}
