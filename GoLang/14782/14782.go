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

	sum := n + 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
