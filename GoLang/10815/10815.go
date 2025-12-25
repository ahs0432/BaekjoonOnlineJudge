package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func search(arr []int64, s, e int, num int64) int {
	for s <= e {
		mid := (s + e) / 2

		if arr[mid] == num {
			return 1
		} else if arr[mid] > num {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	var e int64

	fmt.Fscanln(reader, &n)
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &arr[i])
		} else {
			fmt.Fscan(reader, &arr[i])
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Fscanln(reader, &m)
	for i := 0; i < m; i++ {
		if i == m-1 {
			fmt.Fscanln(reader, &e)
		} else {
			fmt.Fscan(reader, &e)
		}
		fmt.Fprint(writer, search(arr, 0, n-1, e), " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
