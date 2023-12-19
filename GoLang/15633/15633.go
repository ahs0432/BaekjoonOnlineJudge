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

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	fmt.Println(sum*5 - 24)
}
