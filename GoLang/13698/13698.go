package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	cup := []int{1, 2, 3, 4}
	var a, b int

	for _, c := range s {
		if c == 'A' {
			cup[0], cup[1] = cup[1], cup[0]
		} else if c == 'B' {
			cup[0], cup[2] = cup[2], cup[0]
		} else if c == 'C' {
			cup[0], cup[3] = cup[3], cup[0]
		} else if c == 'D' {
			cup[1], cup[2] = cup[2], cup[1]
		} else if c == 'E' {
			cup[1], cup[3] = cup[3], cup[1]
		} else if c == 'F' {
			cup[2], cup[3] = cup[3], cup[2]
		}
	}

	for i := 0; i < 4; i++ {
		if cup[i] == 1 {
			a = i + 1
		} else if cup[i] == 4 {
			b = i + 1
		}
	}

	fmt.Println(a)
	fmt.Println(b)
}
