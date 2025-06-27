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

	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 'L' && s[i+1] == 'L' {
			count++
			i++
		}
	}

	if count <= 1 {
		fmt.Println(n)
	} else {
		fmt.Println(n + 1 - count)
	}
}
