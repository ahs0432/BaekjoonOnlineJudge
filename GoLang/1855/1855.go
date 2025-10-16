package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscanln(reader, &k)

	var s string
	fmt.Fscanln(reader, &s)

	table := make([]string, 0)
	for i := 0; i < len(s); i += k {
		table = append(table, s[i:i+k])
	}

	for i := 0; i < len(table); i++ {
		if i%2 != 0 {
			table[i] = reverse(table[i])
		}
	}

	ans := ""
	for i := 0; i < k; i++ {
		for j := 0; j < len(table); j++ {
			ans += string(table[j][i])
		}
	}
	fmt.Println(ans)
}
