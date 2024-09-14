package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, n int
	fmt.Fscanln(reader, &a, &b, &c, &n)

	count := 0
	for i := 0; i < n/a+1; i++ {
		for j := 0; j < n/b+1; j++ {
			for k := 0; k < n/c+1; k++ {
				if a*i+b*j+c*k == n {
					count = 1
				}
			}
		}
	}
	fmt.Println(count)
}
