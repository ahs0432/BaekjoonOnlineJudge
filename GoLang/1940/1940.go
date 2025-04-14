package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	items := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &items[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if items[i]+items[j] == m {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
