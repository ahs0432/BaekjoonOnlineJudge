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

	var ans int
	var s string
	var alphaMap []bool
	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "#" {
			break
		}

		alphaMap = make([]bool, 26)
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				alphaMap[int(c-'A')] = true
			} else if c >= 'a' && c <= 'z' {
				alphaMap[int(c-'a')] = true
			}
		}

		ans = 0
		for _, alpha := range alphaMap {
			if alpha {
				ans++
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
