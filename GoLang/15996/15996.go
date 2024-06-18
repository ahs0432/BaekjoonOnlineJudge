package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a int64
	fmt.Fscanln(reader, &n, &a)

	m := n / a
	count := m

	for m >= a {
		m /= a
		count += m
	}

	fmt.Println(count)
}
