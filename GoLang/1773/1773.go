package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, c int
	fmt.Fscanln(reader, &n, &c)

	var num, tmp int
	table := make([]int, c+1)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &num)
		tmp = num

		if table[num] == 0 {
			for {
				if num > c {
					break
				} else {
					table[num] = 1
					num += tmp
				}
			}
		}
	}

	count := 0
	for i := 1; i <= c; i++ {
		if table[i] != 0 {
			count++
		}
	}
	fmt.Println(count)
}
