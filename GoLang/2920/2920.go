package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	types := 0

	for i := 1; i <= 8; i++ {
		var n int
		fmt.Fscan(reader, &n)

		if i == 1 {
			if n == i {
				types = 1
			} else if n == 8 {
				types = 2
			}
		} else {
			if types == 1 && i != n {
				types = 0
			}

			if types == 2 && 8-(i-1) != n {
				types = 0
			}
		}

		if types == 0 {
			fmt.Println("mixed")
			return
		}
	}

	if types == 1 {
		fmt.Println("ascending")
	} else {
		fmt.Println("descending")
	}
}
