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

	n = (n % 8)
	if n == 0 || n == 2 {
		fmt.Println(2)
	} else if n == 1 {
		fmt.Println(n)
	} else if n == 3 || n == 7 {
		fmt.Println(3)
	} else if n == 4 || n == 6 {
		fmt.Println(4)
	} else if n == 5 {
		fmt.Println(5)
	}
}
