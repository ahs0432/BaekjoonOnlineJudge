package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	ans := n / m
	if n%m != 0 {
		ans++
	}
	fmt.Println(ans)
}
