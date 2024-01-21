package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, r, a int
	fmt.Fscanln(reader, &l, &r, &a)

	max, min := 0, 0
	sum := l + r + a
	if l < r {
		max = r
		min = l
	} else {
		max = l
		min = r
	}

	tmp := a - (max - min)

	if tmp >= 0 {
		fmt.Println(sum - tmp%2)
	} else {
		fmt.Println(sum + (tmp))
	}
}
