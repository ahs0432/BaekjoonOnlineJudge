package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, e, f, g, h, i int
	fmt.Fscanln(reader, &a, &b, &c, &d, &e, &f, &g, &h, &i)

	if a > 100 || b > 100 || c > 200 || d > 200 || e > 300 || f > 300 || g > 400 || h > 400 || i > 500 {
		fmt.Println("hacker")
	} else if a+b+c+d+e+f+g+h+i < 100 {
		fmt.Println("none")
	} else {
		fmt.Println("draw")
	}
}
