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

	var r, e, c int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &r, &e, &c)
		e = e - c

		if r < e {
			fmt.Println("advertise")
		} else if r == e {
			fmt.Println("does not matter")
		} else {
			fmt.Println("do not advertise")
		}
	}
}
