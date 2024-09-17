package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var a, b, gcdv int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		gcdv = gcd(a, b)
		fmt.Fprintln(writer, (a*b)/gcdv, gcdv)
	}
	writer.Flush()
}
