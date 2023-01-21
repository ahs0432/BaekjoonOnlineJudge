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

	if n == 1 {
		fmt.Println("*")
	} else {
		f := ""
		s := " "
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				f += "* "
			} else {
				s += "* "
			}
		}

		for i := 1; i <= n*2; i++ {
			if i%2 == 1 {
				fmt.Println(f)
			} else {
				fmt.Println(s)
			}
		}
	}
}
