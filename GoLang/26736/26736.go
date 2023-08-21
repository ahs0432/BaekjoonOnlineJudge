package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w string
	fmt.Fscanln(reader, &w)

	a := 0
	b := 0

	for _, c := range w {
		if c == 'A' {
			a++
		} else {
			b++
		}
	}

	fmt.Println(a, ":", b)
}
