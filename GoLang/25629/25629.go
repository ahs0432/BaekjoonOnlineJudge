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

	var tmp int
	var odd, even []int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp%2 == 0 {
			even = append(even, tmp)
		} else {
			odd = append(odd, tmp)
		}
	}

	ans := 1
	if len(odd)-len(even) > 1 || len(odd)-len(even) < 0 {
		ans = 0
	}
	fmt.Println(ans)
}
