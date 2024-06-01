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

	ans := 0
	for i := 0; i < 6; i++ {
		for j := i; j < 6; j++ {
			if i+j == n {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
