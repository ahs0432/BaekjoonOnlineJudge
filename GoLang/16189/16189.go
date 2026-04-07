package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	var n int64
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &n)

	check := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			check = false
			break
		}
	}

	if check {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
