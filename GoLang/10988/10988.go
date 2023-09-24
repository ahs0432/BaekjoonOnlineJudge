package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			fmt.Println("0")
			return
		}
	}
	fmt.Println("1")
}
