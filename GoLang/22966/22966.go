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

	ansa := 5

	var a int
	var anss, s string

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s, &a)
		if ansa > a {
			anss = s
			ansa = a
		}
	}
	fmt.Println(anss)
}
