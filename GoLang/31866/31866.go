package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	if a == b {
		fmt.Println("=")
	} else if (a == 0 && b == 5) || (a == 2 && b == 0) || (a == 5 && b == 2) {
		fmt.Println("<")
	} else if (b == 0 && a == 5) || (b == 2 && a == 0) || (b == 5 && a == 2) {
		fmt.Println(">")
	} else if a != 0 && a != 2 && a != 5 {
		if b != 0 && b != 2 && b != 5 {
			fmt.Println("=")
		} else {
			fmt.Println("<")
		}
	} else if b != 0 && b != 2 && b != 5 {
		if a != 0 && a != 2 && a != 5 {
			fmt.Println("=")
		} else {
			fmt.Println(">")
		}
	}

}
