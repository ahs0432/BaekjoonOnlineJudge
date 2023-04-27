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

	if n == 0 {
		fmt.Println("divide by zero")
	} else {
		sum := 0
		for i := 0; i < n; i++ {
			var tmp int
			if i == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			sum += tmp
		}

		if sum == 0 {
			fmt.Println("divide by zero")
		} else {
			fmt.Println("1.00")
		}
	}
}
