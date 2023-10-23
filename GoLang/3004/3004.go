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

	sum := 2
	for i := 2; i <= n; i++ {
		sum += ((i / 2) + 1)
	}

	fmt.Println(sum)
}
