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

	var s string
	for s != "Y" {
		fmt.Println("? 1")
		fmt.Fscanln(reader, &s)
	}
	fmt.Println("! 1")
}
