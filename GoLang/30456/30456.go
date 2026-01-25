package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscanln(reader, &n, &l)

	for i := 0; i < l-1; i++ {
		fmt.Print(1)
	}
	fmt.Println(n)
}
