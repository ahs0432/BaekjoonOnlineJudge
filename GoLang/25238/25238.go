package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b float64
	fmt.Fscanln(reader, &a, &b)

	if a-(a/100)*b < 100 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
