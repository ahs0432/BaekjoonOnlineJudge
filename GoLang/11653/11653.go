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

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			fmt.Println(i)
		}
	}
}
