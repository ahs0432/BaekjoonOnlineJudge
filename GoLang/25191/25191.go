package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a, b int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &a, &b)

	sum := a/2 + b
	if sum < n {
		fmt.Println(sum)
	} else {
		fmt.Println(n)
	}
}
