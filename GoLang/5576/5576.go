package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	w := make([]int, 10)
	k := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &w[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &k[i])
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i] > w[j]
	})

	sort.Slice(k, func(i, j int) bool {
		return k[i] > k[j]
	})

	fmt.Println(w[0]+w[1]+w[2], k[0]+k[1]+k[2])
}
