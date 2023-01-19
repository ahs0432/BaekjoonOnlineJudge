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

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}

		if i == n-1 {
			for j := 0; j < (n*2)-1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		} else {
			if i != 0 {
				fmt.Print("*")
			}
			for j := 0; j < (i-1)*2+1; j++ {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
	}
}
