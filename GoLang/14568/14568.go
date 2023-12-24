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

	sum := 0
	for i := 2; i <= n-2; i += 2 {
		sum += (n - i - 2) / 2
	}
	fmt.Println(sum)
}
