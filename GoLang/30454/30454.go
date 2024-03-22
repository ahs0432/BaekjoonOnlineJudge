package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscanln(reader, &n, &l)

	var s string
	count, max, maxcount := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)

		count = 0
		prev := byte('0')
		for j := 0; j < l; j++ {
			if prev != s[j] {
				prev = s[j]
				if s[j] == '1' {
					count++
				}
			}
		}

		if max < count {
			max = count
			maxcount = 1
		} else if max == count {
			maxcount++
		}
	}
	fmt.Println(max, maxcount)
}
