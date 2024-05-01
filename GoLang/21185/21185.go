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

	if n%4 == 0 {
		fmt.Println("Even")
	} else if n%2 == 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Either")
	}
}
