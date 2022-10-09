package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a int, b int) int {
	for b != 0 {
		n := a % b
		a = b
		b = n
	}

	fmt.Println(a)
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2 int
	fmt.Fscanln(reader, &n1, &n2)

	if n1 < n2 {
		tmp := n1
		n1 = n2
		n2 = tmp
	}

	gcd := gcd(n1, n2)
	fmt.Println(n1 * n2 / gcd)
}
