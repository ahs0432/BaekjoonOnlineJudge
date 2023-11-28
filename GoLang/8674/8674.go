package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y int
	fmt.Fscanln(reader, &x, &y)

	if x%2 == 0 || y%2 == 0 {
		fmt.Println(0)
	} else if x < y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}
