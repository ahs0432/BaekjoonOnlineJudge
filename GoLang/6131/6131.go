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

	result := 0

	for i := 1; i <= 1000; i++ {
		for j := 1 + 1; j <= 1000; j++ {
			if j*j == i*i+n {
				result++
			}
		}
	}

	fmt.Println(result)
}
