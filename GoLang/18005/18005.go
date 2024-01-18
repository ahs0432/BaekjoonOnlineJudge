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

	if n%2 == 1 {
		fmt.Println(0)
	} else if n/2%2 == 0 {
		fmt.Println(2)
	} else {
		fmt.Println(1)
	}
}
