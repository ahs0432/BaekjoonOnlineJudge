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
	if n%2 == 0 {
		fmt.Println("I LOVE CBNU")
	} else {
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()

		n = n/2 + 1
		f := n
		l := n

		for i := 0; i < n; i++ {
			for j := 0; j < f; j++ {
				if j != f-1 {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}

			for j := f; j < l; j++ {
				if j != l-1 {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Println()
			f--
			l++
		}
	}

}
