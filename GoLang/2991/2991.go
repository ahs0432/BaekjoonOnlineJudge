package main

import (
	"bufio"
	"fmt"
	"os"
)

func attack(a, b, c, d, n int) int {
	at := 0
	if 0 < n%(a+b) && n%(a+b) <= a {
		at++
	}

	if 0 < n%(c+d) && n%(c+d) <= c {
		at++
	}

	return at
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, p, m, n int
	fmt.Fscanln(reader, &a, &b, &c, &d)
	fmt.Fscanln(reader, &p, &m, &n)

	fmt.Println(attack(a, b, c, d, p))
	fmt.Println(attack(a, b, c, d, m))
	fmt.Println(attack(a, b, c, d, n))
}
