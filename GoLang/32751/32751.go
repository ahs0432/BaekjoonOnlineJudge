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

	table := make([]int, 4)
	for i := 0; i < 4; i++ {
		if i == 3 {
			fmt.Fscanln(reader, &table[i])
		} else {
			fmt.Fscan(reader, &table[i])
		}
	}

	var s string
	fmt.Fscanln(reader, &s)

	check := true
	if s[0] != 'a' || s[len(s)-1] != 'a' {
		check = false
	} else {
		table[0]--
		for i := 1; i < n; i++ {
			if s[i] == s[i-1] {
				check = false
				break
			}

			table[s[i]-'a']--
			if table[s[i]-'a'] < 0 {
				check = false
				break
			}
		}
	}

	if check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
