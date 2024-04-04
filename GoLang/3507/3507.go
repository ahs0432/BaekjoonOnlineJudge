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

	cnt := 0
	if n <= 198 {
		for i := 1; i < 100; i++ {
			for j := 1; j < 100; j++ {
				if i+j == n {
					cnt += 1
				}
			}
		}
	}
	fmt.Println(cnt)
}
