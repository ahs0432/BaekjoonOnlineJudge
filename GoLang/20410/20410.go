package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, seed, x1, x2 int
	fmt.Fscanln(reader, &m, &seed, &x1, &x2)

	for a := 0; a < 100; a++ {
		for c := 0; c < 100; c++ {
			if x1 == (a*seed+c)%m {
				if x2 == (a*((a*seed+c)%m)+c)%m {
					fmt.Println(a, c)
					return
				}
			}
		}
	}
}
