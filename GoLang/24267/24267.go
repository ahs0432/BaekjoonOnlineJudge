package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)
	fmt.Println((n - 2) * (n - 1) * n / 6)
	fmt.Println(3)
}
