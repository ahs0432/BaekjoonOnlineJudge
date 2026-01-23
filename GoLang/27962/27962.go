package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, count int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	var a, b string
	for i := 1; i <= n; i++ {
		a = s[0:i]
		b = s[n-i:]
		count = 0

		for j := 0; j < i; j++ {
			if a[j] != b[j] {
				count++
			}
		}
		if count == 1 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
