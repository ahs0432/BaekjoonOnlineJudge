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

	n = n / 3
	fmt.Println((n - 1) * (n - 2) / 2)
}
