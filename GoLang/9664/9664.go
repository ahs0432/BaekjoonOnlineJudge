package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, o int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &o)

	p := o / (n - 1)
	if o-p*(n-1) == 0 {
		fmt.Println(o+p-1, o+p)
	} else {
		fmt.Println(o+p, o+p)
	}
}
