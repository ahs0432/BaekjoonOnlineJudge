package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, d int
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &d)

	if m == 1 {
		fmt.Println("Before")
	} else if m == 2 {
		if d < 18 {
			fmt.Println("Before")
		} else if d == 18 {
			fmt.Println("Special")
		} else {
			fmt.Println("After")
		}
	} else {
		fmt.Println("After")
	}
}
