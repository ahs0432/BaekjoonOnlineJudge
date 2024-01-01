package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := make([]int, 3)
	fmt.Fscanln(reader, &t[0], &t[1], &t[2])

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	if t[0] == t[1] && t[1] == t[2] && t[0] == t[2] {
		fmt.Println("2")
	} else if t[2]*t[2] == (t[0]*t[0])+(t[1]*t[1]) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
