package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	if len(n) < len(m) {
		fmt.Println("no")
	} else {
		fmt.Println("go")
	}
}
