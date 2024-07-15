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

	var a, b string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	ans := 0
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
