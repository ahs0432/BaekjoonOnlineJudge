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

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanln(reader, &s)
		r := []rune(s)

		for j, c := range r {
			if c == 'Z' {
				r[j] = 'A'
			} else {
				r[j] = rune(c + 1)
			}
		}
		fmt.Print("String #")
		fmt.Println(i + 1)
		fmt.Println(string(r))
		fmt.Println()
	}
}
