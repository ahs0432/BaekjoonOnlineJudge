package main

import (
	"bufio"
	"fmt"
	"os"
)

func bs(minsu []int, now int) int {
	low := 0
	left := 0
	right := len(minsu)

	for left+1 < right {
		mid := (left + right) / 2

		if minsu[mid] < now {
			left = mid
		} else if minsu[mid] > now {
			right = mid
			low = mid
		} else {
			if mid+1 != len(minsu) {
				low = mid + 1
			}
			break
		}
	}

	return low
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m, &k)

	table := make([]int, n+1)
	minsu := []int{}
	check := make([]bool, m)

	for i := 0; i < m; i++ {
		var tmp int
		if i == m-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}

		table[tmp] = 1
	}

	for i := 1; i <= n; i++ {
		if table[i] != 0 {
			minsu = append(minsu, i)
		}
	}

	for i := 0; i < k; i++ {
		var now int
		if i == k-1 {
			fmt.Fscanln(reader, &now)
		} else {
			fmt.Fscan(reader, &now)
		}

		num := 0
		if !(now < minsu[0] || now >= minsu[len(minsu)-1]) {
			num = bs(minsu, now)
		}

		for j := num; j < m && check[num]; j++ {
			num++
		}

		fmt.Fprintln(writer, minsu[num])
		check[num] = true
	}
	writer.Flush()
}
