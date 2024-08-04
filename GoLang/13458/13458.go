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

	a := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	var b, c int
	fmt.Fscanln(reader, &b, &c)

	ans := int64(n)

	for i := 0; i < n; i++ {
		a[i] -= b
		if a[i] < 1 {
			continue
		}

		ans += int64(a[i] / c)
		if a[i]%c != 0 {
			ans++
		}
	}
	fmt.Println(ans)

}
