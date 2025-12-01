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

	count := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
