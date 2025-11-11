package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d int
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &c, &d)

	h := a + c
	y := b + d

	if h < y {
		fmt.Println("Hanyang Univ.")
	} else if y < h {
		fmt.Println("Yongdap")
	} else {
		fmt.Println("Either")
	}
}
