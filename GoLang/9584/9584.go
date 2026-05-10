package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	fmt.Fscanln(reader, &s)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	var check bool
	ans := []string{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		check = true
		for j := range s {
			if s[j] != tmp[j] && s[j] != '*' {
				check = false
				break
			}
		}

		if check {
			ans = append(ans, tmp)
		}
	}
	fmt.Fprintln(writer, len(ans))
	for i := range ans {
		fmt.Fprintln(writer, ans[i])
	}
	writer.Flush()
}
