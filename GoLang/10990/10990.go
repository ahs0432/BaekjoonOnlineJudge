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

	for i := 1; i <= n; i++ {
		for j := i; j < n; j++ {
			fmt.Print(" ")
		}

		if i == 1 {
			fmt.Println("*")
		} else {
			fmt.Print("*")
			for j := 0; j < 1+((i-2)*2); j++ {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
	}
}
