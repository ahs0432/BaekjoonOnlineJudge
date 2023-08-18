package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := 0
	for i := 0; i < 4; i++ {
		var p string
		var h int
		fmt.Fscanln(reader, &p, &h)

		if p == "Es" {
			res += h * 21
		} else {
			res += h * 17
		}
	}
	fmt.Println(res)
}
