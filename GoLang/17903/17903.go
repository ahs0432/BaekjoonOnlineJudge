package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var tmp int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if m-1 == j {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
		}
	}

	if n >= 8 {
		fmt.Println("satisfactory")
	} else {
		fmt.Println("unsatisfactory")
	}
}
