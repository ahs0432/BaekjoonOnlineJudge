package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y int
	fmt.Fscanln(reader, &x, &y)

	count := 1

	for {
		if y*count-x*count >= y {
			break
		}
		count++
	}
	fmt.Println(count)
}
