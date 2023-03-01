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

	sum := n / 5
	if n%5 != 0 {
		sum += 1
	}
	fmt.Println(sum)
}
