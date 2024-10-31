package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, x, y int
	fmt.Fscanln(reader, &a, &b, &c, &x, &y)

	if a+b < 2*c {
		fmt.Println(a*x + b*y)
	} else {
		fmt.Println(2*c*min(x, y) + min(a, 2*c)*max(0, x-y) + min(b, 2*c)*max(0, y-x))
	}
}
