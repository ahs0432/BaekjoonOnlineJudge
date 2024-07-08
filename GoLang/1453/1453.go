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

	ans := 0
	var tmp int
	table := make([]bool, 100+1)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if table[tmp] {
			ans++
		} else {
			table[tmp] = true
		}
	}
	fmt.Println(ans)
}
