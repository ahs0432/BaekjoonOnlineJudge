package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	tmp := n
	var check, ans int
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			if tmp-i < 0 {
				ans = abs(tmp - i)
				check = 1
			}
		} else {
			if tmp-i < 0 {
				ans = 0
				check = 1
			}
		}
		tmp -= i
		if check == 1 {
			break
		}
	}
	fmt.Println(ans)
}
