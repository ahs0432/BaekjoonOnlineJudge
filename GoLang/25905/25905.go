package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]float64, 10)

	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] > table[j]
	})

	ans := 1.0
	for i := 0; i < 9; i++ {
		ans *= table[i] / float64(i+1)
	}

	fmt.Printf("%.06f\n", ans*math.Pow(10, 9))
}
