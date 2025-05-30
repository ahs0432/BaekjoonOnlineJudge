package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	if n == 1 || m == 1 {
		fmt.Println(0)
	} else {
		fmt.Println((n - 1) * (m - 1) * 2)
	}
}
