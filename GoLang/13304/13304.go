package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	st := []int{0, 0, 0, 0, 0}

	var s, y int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s, &y)

		if y == 1 || y == 2 {
			st[0]++
		} else if s == 0 {
			if y == 3 || y == 4 {
				st[1]++
			} else if y == 5 || y == 6 {
				st[2]++
			}
		} else {
			if y == 3 || y == 4 {
				st[3]++
			} else if y == 5 || y == 6 {
				st[4]++
			}
		}
	}

	ans := 0
	for _, s := range st {
		ans += s / k
		if s%k != 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
