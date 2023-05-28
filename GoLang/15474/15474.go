package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a, b, c, d int
	fmt.Fscanln(reader, &n, &a, &b, &c, &d)

	ar := n / a
	if n%a != 0 {
		ar++
	}
	ar = ar * b

	br := n / c
	if n%c != 0 {
		br++
	}
	br = br * d

	if ar < br {
		fmt.Println(ar)
	} else {
		fmt.Println(br)
	}
}
