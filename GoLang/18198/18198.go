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

	a := 0
	b := 0

	for i := 0; i < len(n); i += 2 {
		if n[i] == 'A' {
			a += int(n[i+1]) - 48
		} else {
			b += int(n[i+1]) - 48
		}
	}

	if a < b {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
