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

	var a, b int
	fmt.Fscanln(reader, &a, &b)
	fmt.Println(gcd(a, b))
}
