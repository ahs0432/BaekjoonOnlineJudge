package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	odd := []int{}
	table := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &table[i])

		if table[i]%2 != 0 {
			odd = append(odd, table[i])
		}
	}

	ans := 1
	if len(odd) != 0 {
		for _, o := range odd {
			ans *= o
		}
	} else {
		for _, t := range table {
			ans *= t
		}
	}

	fmt.Println(ans)
}
