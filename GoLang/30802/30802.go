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

	s := make([]int, 6)
	for i := 0; i < 6; i++ {
		if i == 5 {
			fmt.Fscanln(reader, &s[i])
		} else {
			fmt.Fscan(reader, &s[i])
		}
	}

	var t, p int
	fmt.Fscanln(reader, &t, &p)

	ans := 0
	for i := 0; i < 6; i++ {
		ans += (s[i] + t - 1) / t
	}
	fmt.Println(ans)
	fmt.Println(n/p, n%p)
}
