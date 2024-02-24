package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, ans := 0.0, 0
	var c, w float64
	snack := []string{"S", "N", "U"}

	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &c, &w)
		c, w = c*10, w*10

		if c >= 5000 {
			c -= 500
		}

		if (w / c) > res {
			res = w / c
			ans = i
		}
	}
	fmt.Println(snack[ans])
}
