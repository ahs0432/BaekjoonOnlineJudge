package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	for _, c := range n {
		if c >= 65 && c <= 90 {
			fmt.Print(string(c + 32))
		} else {
			fmt.Print(string(c - 32))
		}
	}
	fmt.Println()
}
