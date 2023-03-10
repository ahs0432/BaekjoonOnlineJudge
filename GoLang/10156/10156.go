package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, n, m int
	fmt.Fscanln(reader, &k, &n, &m)

	d := (k * n) - m
	if d <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(d)
	}
}
