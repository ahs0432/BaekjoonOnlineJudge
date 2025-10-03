package main

import (
	"bufio"
	"fmt"
	"os"
)

var num = []string{"000000", "001111", "010011", "011100", "100110", "101001", "110101", "111010"}

func check(s string, i int) int {
	ans := 0
	for j := 0; j < 6; j++ {
		if num[i][j] != s[j] {
			ans++
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s, ans string
	fmt.Fscanln(reader, &s)

	var c bool
	var tmp string
	for i := 0; i < n; i++ {
		tmp = s[i*6 : i*6+6]
		c = false

		for j := 0; j < 8; j++ {
			if check(tmp, j) <= 1 {
				c = true
				ans += string(byte('A' + j))
				break
			}
		}

		if !c {
			fmt.Println(i + 1)
			return
		}
	}

	fmt.Println(ans)
}
