package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, p, count int
	fmt.Fscanln(reader, &n, &p)

	tmp := n
	table := make([]int, 98)
	for {
		tmp = (tmp * n) % p
		if table[tmp] == 2 {
			break
		}
		table[tmp]++
	}

	for i := 0; i < p; i++ {
		if table[i] == 2 {
			count++
		}
	}
	fmt.Println(count)
}
