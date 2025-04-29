package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(a, m, p int64) int64 {
	r := int64(1)
	for m > 0 {
		if (m & 1) == 1 {
			r = r * a % p
		}
		a = a * a % p
		m >>= 1
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var a, n, p int64
	var ans int64
	for i := 1; i <= t; i++ {
		fmt.Fscanln(reader, &a, &n, &p)

		ans = a % p
		for j := int64(2); j <= n; j++ {
			ans = calc(ans, j, p)
		}

		fmt.Fprintf(writer, "Case #%d: %v\n", i, ans)
	}

	writer.Flush()
}
