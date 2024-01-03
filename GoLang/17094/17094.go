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
	for _, c := range s {
		if c == 'e' {
			count++
		} else {
			count--
		}
	}

	if count == 0 {
		fmt.Println("yee")
	} else if count > 0 {
		fmt.Println("e")
	} else {
		fmt.Println("2")
	}
}
