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

	if n <= 10000 {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Time limit exceeded")
	}
}
