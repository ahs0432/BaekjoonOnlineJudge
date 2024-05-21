package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, n, m int
	fmt.Fscanln(reader, &s, &n, &m)

	count := 0
	var num int
	for i := 0; i < n+m; i++ {
		fmt.Fscanln(reader, &num)

		if num == 1 {
			if s == count {
				s += s
			}
			count++
		} else {
			count--
		}
	}
	fmt.Println(s)
}
