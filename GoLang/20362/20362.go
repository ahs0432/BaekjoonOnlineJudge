package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n, &s)

	name := make([]string, n)
	chat := make([]string, n)

	check := -1
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &name[i], &chat[i])
		if name[i] == s {
			check = i
		}
	}

	ans := 0
	for i := 0; i < check; i++ {
		if chat[check] == chat[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
