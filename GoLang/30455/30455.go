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
		fmt.Println("Goose")
	} else {
		fmt.Println("Duck")
	}
}
