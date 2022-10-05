package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, w, h int
	fmt.Fscanln(reader, &x, &y, &w, &h)

	min := 2147483647

	if min > x {
		min = x
	}

	if min > y {
		min = y
	}

	if min > (w - x) {
		min = (w - x)
	}

	if min > (h - y) {
		min = (h - y)
	}

	fmt.Println(min)
}
