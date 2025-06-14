package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, n int
	fmt.Fscanln(reader, &a, &b, &n)

	ans := 1001
	var cost, cycle int
	var table []int
	var d bool
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &cost, &cycle)

		d = false
		table = make([]int, cycle)
		for j := 0; j < cycle; j++ {
			if j == cycle-1 {
				fmt.Fscanln(reader, &table[j])
			} else {
				fmt.Fscan(reader, &table[j])
			}
		}

		for j := 0; j < cycle; j++ {
			if a == table[j] {
				d = true
			} else if b == table[j] {
				if !d {
					break
				} else {
					if cost < ans {
						ans = cost
					}
				}
			}
		}
	}

	if ans != 1001 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}
