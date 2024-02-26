package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func operation(a, b int, o string) int {
	if o == "+" {
		return a + b
	} else if o == "-" {
		return a - b
	} else if o == "*" {
		return a * b
	} else {
		return a / b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	k := make([]int, 3)
	o := make([]string, 2)

	fmt.Fscanln(reader, &k[0], &o[0], &k[1], &o[1], &k[2])
	res := []int{operation(operation(k[0], k[1], o[0]), k[2], o[1]), operation(k[0], operation(k[1], k[2], o[1]), o[0])}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	fmt.Println(res[0])
	fmt.Println(res[1])
}
