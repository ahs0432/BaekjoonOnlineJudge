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

	arr := []int{}
	table := make([]int, n*2)

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		arr = append(arr, tmp)
	}

	for i := 0; i < n; i++ {
		arr = append(arr, arr[i])
	}

	table[0] = arr[0]

	for i := 1; i < k; i++ {
		table[i] = table[i-1] + arr[i]
	}

	var ans int
	for i := k; i < n*2; i++ {
		table[i] = table[i-1] - arr[i-k] + arr[i]
		if ans < table[i] {
			ans = table[i]
		}
	}
	fmt.Println(ans)
}
