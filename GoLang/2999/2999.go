package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	x, y := 0, 0
	for i := 1; i < len(s)/2+1; i++ {
		for j := i; j < len(s)+1; j++ {
			if i*j == len(s) {
				x, y = i, j
			}
		}
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(string(s[i+j*x]))
		}
	}
	fmt.Println()
}
