package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, l, r int
	fmt.Fscanln(reader, &x, &l, &r)

	if x < l {
		fmt.Println(l)
	} else if x > r {
		fmt.Println(r)
	} else {
		fmt.Println(x)
	}
}
