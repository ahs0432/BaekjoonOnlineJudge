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

	var a, b, tmp int
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		a += tmp
	}

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		b += tmp
	}

	fmt.Println(b, a)
}
