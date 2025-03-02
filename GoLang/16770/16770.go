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

	bk := make([]int, 1001)

	var a, b, v int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &v)
		bk[a] += v
		bk[b] -= v
	}

	ans, count := 0, 0
	for i := 0; i <= 1000; i++ {
		count += bk[i]
		if ans < count {
			ans = count
		}
	}

	fmt.Println(ans)
}
