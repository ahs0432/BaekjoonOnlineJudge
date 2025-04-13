package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	ans := []string{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if len(tmp) == 3 {
			ans = append(ans, tmp)
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	fmt.Println(ans[0])
}
