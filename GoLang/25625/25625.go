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

	if x > y {
		fmt.Println(x + y)
	} else {
		fmt.Println(y - x)
	}
}
