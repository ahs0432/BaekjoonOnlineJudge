package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a float64
	fmt.Fscanln(reader, &a)

	var m, n float64
	fmt.Fscanln(reader, &m, &n)

	if m > n {
		m, n = n, m
	}

	ans := max(m, n/a)
	ans = min((m/a)*2, ans)
	fmt.Printf("%.7f\n", ans)
}
