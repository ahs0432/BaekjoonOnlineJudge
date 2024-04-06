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

	if n < 206 {
		fmt.Println(1)
	} else if n < 218 {
		fmt.Println(2)
	} else if n < 229 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
