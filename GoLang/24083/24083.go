package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	if (a+b)%12 == 0 {
		fmt.Println(12)
	} else {
		fmt.Println((a + b) % 12)
	}
}
