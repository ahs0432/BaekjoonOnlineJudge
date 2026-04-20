package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, c int
	fmt.Fscanln(reader, &r, &c)
	if r > c {
		fmt.Println((c-1)*2 + 1)
	} else {
		fmt.Println((r - 1) * 2)
	}
}
