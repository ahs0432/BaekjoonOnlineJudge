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

	costs := make([]int, 1001)
	jin := 0
	cnt := 0

	var region string
	var cost int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &region, &cost)

		if region == "jinju" {
			jin = cost
		} else {
			costs[cost] += 1
		}
	}

	for i := jin + 1; i < 1001; i++ {
		cnt += costs[i]
	}
	fmt.Println(jin)
	fmt.Println(cnt)
}
