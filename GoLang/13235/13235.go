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

	check := true

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			check = false
			break
		}
	}

	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
