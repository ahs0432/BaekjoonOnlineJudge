package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int64
	fmt.Fscanln(reader, &n, &m)

	var res int64 = 1

	if m <= n {
		res = 0
	} else {
		for i := n; i > 1; i-- {
			res *= i
			res %= m

			if res == 0 {
				break
			}
		}
	}

	fmt.Println(res)
}
