package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	a, b := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			a++
		} else {
			b++
		}
	}

	if a > b {
		fmt.Println("A")
	} else if a < b {
		fmt.Println("B")
	} else {
		fmt.Println("Tie")
	}
}
