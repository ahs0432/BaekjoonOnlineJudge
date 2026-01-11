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

	var check bool
	var tmp, prev int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp == prev {
			fmt.Println("NO")
			return
		}

		if tmp > prev {
			if check {
				fmt.Println("NO")
				return
			} else {
				prev = tmp
			}
		} else {
			if !check {
				check = true
			}
			prev = tmp
		}
	}

	fmt.Println("YES")
}
