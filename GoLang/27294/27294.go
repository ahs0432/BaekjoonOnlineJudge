package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	if m == 1 || (n < 12 || n > 16) {
		fmt.Println("280")
	} else {
		fmt.Println("320")
	}
}
