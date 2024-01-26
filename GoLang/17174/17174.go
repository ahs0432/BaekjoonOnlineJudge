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

	a := 0
	for n != 0 {
		a += n
		n = n / m
	}
	fmt.Println(a)
}
