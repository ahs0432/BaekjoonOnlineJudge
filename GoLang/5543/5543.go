package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := []int{2147483647, 2147483647}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3-i; j++ {
			var tmp int
			fmt.Fscanln(reader, &tmp)
			if t[i] > tmp {
				t[i] = tmp
			}
		}
	}

	fmt.Println(t[0] + t[1] - 50)
}
