package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w, b int
	fmt.Fscanln(reader, &w, &b)

	if w >= b {
		fmt.Println(b)
	} else {
		fmt.Println(w + 1)
	}
}
