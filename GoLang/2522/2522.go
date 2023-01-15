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

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
