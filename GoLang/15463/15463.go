package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var table [2001][2001]int

	var a, b, c, d int
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &a, &b, &c, &d)
		a += 1000
		b += 1000
		c += 1000
		d += 1000

		for j := a; j < c; j++ {
			for k := b; k < d; k++ {
				if i != 2 {
					table[j][k] = 1
				} else {
					table[j][k] = 0
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 2001; i++ {
		for j := 0; j < 2001; j++ {
			ans += table[i][j]
		}
	}
	fmt.Println(ans)
}
