package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	for i := 0; i < m; i++ {
		if n%2 == 0 {
			n = (n / 2) ^ 6
		} else {
			n = (n * 2) ^ 6
		}
	}
	fmt.Println(n)
}
