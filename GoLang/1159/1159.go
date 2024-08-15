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

	players := make([]byte, 26)
	var tmp string

	ans := []byte{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		players[tmp[0]-97]++

		if players[tmp[0]-97] == 5 {
			ans = append(ans, byte(tmp[0]))
		}
	}

	if len(ans) == 0 {
		fmt.Println("PREDAJA")
	} else {
		sort.Slice(ans, func(i, j int) bool {
			return ans[i] < ans[j]
		})
		fmt.Println(string(ans))
	}
}
