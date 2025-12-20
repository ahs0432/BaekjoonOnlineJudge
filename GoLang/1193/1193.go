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

	i := 1
	for ; i < n; i++ {
		n -= i
	}

	var a, b int
	if i%2 == 0 {
		a = n
		b = i - n + 1
	} else {
		a = i - n + 1
		b = n
	}
	fmt.Printf("%d/%d\n", a, b)
}
