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

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	res := true
	var max, tmp int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &tmp)
		if table[i]+tmp > max {
			max = table[i] + tmp
		}

		if max < table[i+1] {
			res = false
		}
	}

	if res || n == 1 {
		fmt.Println("권병장님, 중대장님이 찾으십니다")
	} else {
		fmt.Println("엄마 나 전역 늦어질 것 같아")
	}
}
