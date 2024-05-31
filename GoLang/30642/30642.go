package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	var m string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &k)

	if m == "annyong" {
		if k%2 == 1 {
			fmt.Println(k)
		} else {
			fmt.Println(k - 1)
		}
	} else {
		if k%2 == 0 {
			fmt.Println(k)
		} else {
			if k-1 != 0 {
				fmt.Println(k - 1)
			} else {
				fmt.Println(k + 1)
			}
		}
	}
}
