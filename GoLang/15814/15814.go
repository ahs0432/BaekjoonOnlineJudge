package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	var n int
	fmt.Fscanln(reader, &n)

	now := []byte(s)
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		now[a], now[b] = now[b], now[a]
	}
	fmt.Println(string(now))
}
