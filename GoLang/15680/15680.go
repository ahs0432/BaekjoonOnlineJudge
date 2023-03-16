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

	if n == 0 {
		fmt.Println("YONSEI")
	} else if n == 1 {
		fmt.Println("Leading the Way to the Future")
	}
}
