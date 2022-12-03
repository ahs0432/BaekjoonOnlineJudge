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

	var t int
	fmt.Fscanln(reader, &t)
	virus := make([]bool, n)
	virus[0] = true

	web := map[int][]int{}

	for i := 0; i < t; i++ {
		var n1, n2 int
		fmt.Fscanln(reader, &n1, &n2)
		n1--
		n2--

		if _, e := web[n1]; !e {
			web[n1] = make([]int, 0)
		}
		web[n1] = append(web[n1], n2)

		if _, e := web[n2]; !e {
			web[n2] = make([]int, 0)
		}
		web[n2] = append(web[n2], n1)
	}

	q := make([]int, 1)
	total := 0

	for len(q) != 0 {
		for _, i := range web[q[0]] {
			if !virus[i] {
				total++
				virus[i] = true
				q = append(q, i)
			}
		}
		q = q[1:]
	}

	fmt.Println(total)
}
