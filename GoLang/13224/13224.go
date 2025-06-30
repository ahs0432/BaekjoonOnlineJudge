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

	cups := []int{1, 0, 0}

	for _, c := range s {
		if c == 'A' {
			cups[0], cups[1] = cups[1], cups[0]
		} else if c == 'B' {
			cups[1], cups[2] = cups[2], cups[1]
		} else if c == 'C' {
			cups[0], cups[2] = cups[2], cups[0]
		}
	}

	for i := 0; i < 3; i++ {
		if cups[i] == 1 {
			fmt.Println(i + 1)
			break
		}
	}
}
