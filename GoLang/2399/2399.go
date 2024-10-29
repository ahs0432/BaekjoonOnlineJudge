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

	table := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	ans := int64(0)
	tmp := int64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp = (table[i] - table[j])
			if tmp < 0 {
				ans += -tmp
			} else {
				ans += tmp
			}
		}
	}
	fmt.Println(ans)
}
