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

	count := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n + 1
		}
		count += 1
	}
	fmt.Println(count)
}
