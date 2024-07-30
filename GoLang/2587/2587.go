package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	sli := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &sli[i])
		sum += sli[i]
	}

	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})

	fmt.Println(sum / 5)
	fmt.Println(sli[2])
}
