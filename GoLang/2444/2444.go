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
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print(" ")
		}
		if i != 1 {
			a += "**"
		}
		fmt.Println(a)
	}

	for i := 2; i <= n; i++ {
		a = a[:len(a)-2]
		a = " " + a

		fmt.Println(a)
	}
}
