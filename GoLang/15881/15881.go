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
	var s string
	fmt.Fscanln(reader, &s)

	word := "pPAp"
	count := 0

	for i := 0; i < n-3; i++ {
		if s[i:i+4] == word {
			count++
			i += 3
		}
	}
	fmt.Println(count)
}
