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

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		count := 0
		for _, s := range tmp {
			if string(s) == "(" {
				count++
			} else {
				count--
			}

			if count < 0 {
				break
			}
		}

		if count == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
