package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var s string
	var count []int
	var max, ans, check int
	for i := 0; i < t; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		count = make([]int, 26)
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				count[c-'a']++
			}
		}

		max = -1
		for _, c := range count {
			if c > max {
				max = c
			}
		}

		ans = 0
		check = 0
		for j, c := range count {
			if max == c {
				ans = j
				check++
			}
		}

		if check > 1 {
			fmt.Fprintln(writer, "?")
		} else {
			fmt.Fprintln(writer, string(rune(ans+'a')))
		}
	}
	writer.Flush()
}
