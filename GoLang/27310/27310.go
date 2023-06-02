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

	sum := len(n)
	for _, c := range n {
		if c == ':' {
			sum++
		}

		if c == '_' {
			sum += 5
		}
	}
	fmt.Println(sum)
}
