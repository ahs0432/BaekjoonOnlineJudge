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

	a := "*"
	for i := 0; i < n-1; i++ {
		a += "**"
	}

	for i := 1; i <= n; i++ {
		if i != 1 {
			a = a[:len(a)-2]
			a = " " + a
		}
		fmt.Println(a)
	}
}
