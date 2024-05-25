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
		fmt.Println(1)
	} else if n == 1 {
		fmt.Println(0)
	} else if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			fmt.Print(8)
		}
		fmt.Println()
	} else {
		fmt.Print(4)
		for i := 0; i < n/2; i++ {
			fmt.Print(8)
		}
		fmt.Println()
	}
}
