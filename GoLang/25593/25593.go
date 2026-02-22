package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	count := 0
	var s string

	table := make(map[string]int)
	workTime := make([]int, 0)
	increments := []int{4, 6, 4, 10}

	for i := 0; i < n; i++ {
		for _, inc := range increments {
			for j := 0; j < 7; j++ {
				if j == 6 {
					fmt.Fscanln(reader, &s)
				} else {
					fmt.Fscan(reader, &s)
				}

				if s == "-" {
					continue
				}

				if _, exists := table[s]; !exists {
					table[s] = count
					workTime = append(workTime, 0)
					count++
				}
				workTime[table[s]] += inc
			}
		}
	}

	if count == 0 {
		fmt.Println("Yes")
		return
	}

	max := -1000000007
	min := 1000000007

	for _, t := range workTime {
		if t > max {
			max = t
		}
		if t < min {
			min = t
		}
	}

	if max-min > 12 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
