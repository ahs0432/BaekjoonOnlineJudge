package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	if a <= b {
		fmt.Println((b - a + 1) * (b + a) / 2)
	} else {
		fmt.Println((a - b + 1) * (b + a) / 2)
	}
}
