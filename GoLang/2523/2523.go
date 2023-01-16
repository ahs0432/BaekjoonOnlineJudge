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

	a := ""
	for i := 0; i < n; i++ {
		a += "*"
		fmt.Println(a)
	}

	for i := 1; i < n; i++ {
		a = a[1:]
		fmt.Println(a)
	}
}
