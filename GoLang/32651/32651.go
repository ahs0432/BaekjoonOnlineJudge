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

	if n%2024 == 0 && n <= 100000 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
