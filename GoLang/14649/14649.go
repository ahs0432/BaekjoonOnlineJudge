package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p, n int
	fmt.Fscanln(reader, &p)
	fmt.Fscanln(reader, &n)

	st := make([]int, 100)

	var num int
	var d string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &num, &d)

		if d == "R" {
			for j := num; j < len(st); j++ {
				st[j]++
			}
		} else {
			for j := 0; j < num-1; j++ {
				st[j]++
			}
		}
	}

	r, g, b := 0, 0, 0

	for i := 0; i < len(st); i++ {
		if st[i]%3 == 0 {
			b++
		} else if st[i]%3 == 1 {
			r++
		} else {
			g++
		}
	}

	fmt.Printf("%.2f\n%.2f\n%.2f\n", float64(p)*float64(b)/100.0, float64(p)*float64(r)/100.0, float64(p)*float64(g)/100.0)
}
