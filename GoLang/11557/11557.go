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
		uni := ""
		max := 0

		var m int
		fmt.Fscanln(reader, &m)

		for j := 0; j < m; j++ {
			var tmp1 string
			var tmp2 int

			fmt.Fscanln(reader, &tmp1, &tmp2)

			if tmp2 > max {
				uni = tmp1
				max = tmp2
			}
		}

		fmt.Println(uni)
	}
}
