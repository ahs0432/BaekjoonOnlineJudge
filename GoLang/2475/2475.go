package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	stored := 0
	for i := 0; i < 5; i++ {
		var n int
		fmt.Fscan(reader, &n)
		stored += (n * n)
	}

	fmt.Println(stored % 10)
}
