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

	var n int
	fmt.Fscanln(reader, &n)

	table := map[int]int{}
	keys := make([]int, n)

	total := 0
	many := []int{}
	manyCount := 0

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &keys[i])

		table[keys[i]]++
		total += keys[i]

		if len(many) == 0 {
			many = []int{keys[i]}
			manyCount = 1
		} else if manyCount < table[keys[i]] {
			many = []int{keys[i]}
			manyCount = table[keys[i]]
		} else if manyCount == table[keys[i]] {
			many = append(many, keys[i])
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Println(int(math.Round(float64(total) / float64(n))))
	fmt.Println(keys[len(keys)/2])

	if len(many) == 1 {
		fmt.Println(many[0])
	} else {
		sort.Slice(many, func(i, j int) bool {
			return many[i] < many[j]
		})
		fmt.Println(many[1])
	}

	fmt.Println(keys[len(keys)-1] - keys[0])
}
