package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l int
	fmt.Fscanln(reader, &l)

	var s string
	fmt.Fscanln(reader, &s)

	table := make([]int64, 4)
	for i := 0; i < l; i++ {
		if s[i] == 'A' {
			table[0]++
		} else if s[i] == 'C' {
			table[1]++
		} else if s[i] == 'G' {
			table[2]++
		} else if s[i] == 'T' {
			table[3]++
		}
	}

	ans := int64(1)
	for i := 0; i < 4; i++ {
		ans = (ans * table[i]) % 1000000007
	}
	fmt.Println(ans)
}
