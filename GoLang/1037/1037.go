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

	divi := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &divi[i])
	}

	sort.Slice(divi, func(i, j int) bool {
		return divi[i] < divi[j]
	})
	fmt.Println(divi[0] * divi[len(divi)-1])
}
