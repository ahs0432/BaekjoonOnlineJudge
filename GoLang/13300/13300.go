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

	var s, y int
	peoples := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s, &y)
		peoples[s][y-1]++
	}

	ans := 0
	for _, people := range peoples {
		for _, count := range people {
			if count != 0 {
				ans += count / k

				if count%k != 0 {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
