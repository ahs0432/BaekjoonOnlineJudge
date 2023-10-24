package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := make([]int, 4)
	fmt.Fscanln(reader, &t[0], &t[1], &t[2], &t[3])

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	fmt.Println(t[0] * t[2])
}
