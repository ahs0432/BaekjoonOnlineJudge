package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	ans := 0
	total := 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		total += tmp
		if total < 0 {
			total = 0
		}
		if total >= m {
			ans++
		}
	}
	fmt.Println(ans)
}
