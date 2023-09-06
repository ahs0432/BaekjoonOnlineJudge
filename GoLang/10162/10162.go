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

	if n%10 != 0 {
		fmt.Println("-1")
	} else {
		a := n / 300
		n %= 300
		b := n / 60
		n %= 60
		c := n / 10

		fmt.Println(a, b, c)
	}
}
