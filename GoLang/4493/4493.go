package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, p1a, p2a int
	fmt.Fscanln(reader, &t)

	var p1, p2 string
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		p1a, p2a = 0, 0
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &p1, &p2)

			if p1 == "R" && p2 == "S" {
				p1a++
			} else if p1 == "S" && p2 == "P" {
				p1a++
			} else if p1 == "P" && p2 == "R" {
				p1a++
			} else if p1 == "S" && p2 == "R" {
				p2a++
			} else if p1 == "P" && p2 == "S" {
				p2a++
			} else if p1 == "R" && p2 == "P" {
				p2a++
			}
		}

		if p1a > p2a {
			fmt.Fprintln(writer, "Player 1")
		} else if p1a < p2a {
			fmt.Fprintln(writer, "Player 2")
		} else {
			fmt.Fprintln(writer, "TIE")
		}
	}

	writer.Flush()
}
