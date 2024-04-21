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

	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp[i])
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	sum := 0
	for i := 0; i < len(tmp)-1; i++ {
		sum += tmp[i]
	}
	fmt.Println(sum)
}
