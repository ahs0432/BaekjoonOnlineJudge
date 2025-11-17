package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	var num, x, y, ans int
	for i := 0; i < 4; i++ {
		fmt.Fscanln(reader, &s)

		for j := 0; j < 4; j++ {
			if s[j] == '.' {
				continue
			}

			num = int(s[j] - 'A')
			x = num / 4
			y = num % 4

			ans += abs(i-x) + abs(j-y)
		}
	}
	fmt.Println(ans)
}
