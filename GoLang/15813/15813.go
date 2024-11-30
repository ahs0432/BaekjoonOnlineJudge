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

	ans := 0
	for _, c := range s {
		ans += int(c - 64)
	}
	fmt.Println(ans)
}
