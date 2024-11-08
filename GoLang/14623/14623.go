package main

import (
	"bufio"
	"fmt"
	"os"
)

func itob(s string) int {
	sum := 0
	now := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			sum += now
		}
		now *= 2
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	ab, bb := itob(a), itob(b)
	fmt.Printf("%b\n", ab*bb)
}
