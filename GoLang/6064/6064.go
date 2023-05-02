package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(m, n int) int {
	r := 0

	for n != 0 {
		r = m % n
		m = n
		n = r
	}

	return m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var m, n, x, y int
		fmt.Fscanln(reader, &m, &n, &x, &y)

		lcm := m * n / gcd(m, n)
		answer := 0

		for {
			if x > lcm {
				answer = -1
				break
			}
			if (x-1)%n+1 == y {
				answer = x
				break
			}
			x += m
		}

		fmt.Fprintln(writer, answer)
	}
	writer.Flush()
}
