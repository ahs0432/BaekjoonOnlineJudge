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

	var s string
	fmt.Fscanln(reader, &s)

	visited := make(map[[2]int]bool)

	x, y := 0, 0
	visited[[2]int{0, 0}] = true

	for _, c := range s {
		switch c {
		case 'E':
			x++
		case 'W':
			x--
		case 'N':
			y++
		case 'S':
			y--
		}
		visited[[2]int{x, y}] = true
	}

	fmt.Println(len(visited))
}
