package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, l int
	fmt.Fscanln(reader, &n, &m, &l)

	table := make([]int, n)
	table[0] = 1
	count := 0

	for i := 0; table[i] != m; {
		if table[i]%2 == 1 {
			i = (i + l) % n
		} else {
			i = (n + (i - l)) % n
		}
		table[i]++
		count++
	}

	fmt.Println(count)
}
