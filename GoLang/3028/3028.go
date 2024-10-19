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

	t := []int{1, 0, 0}
	for _, c := range s {
		if c == 'A' {
			t[0], t[1] = t[1], t[0]
		} else if c == 'B' {
			t[1], t[2] = t[2], t[1]
		} else if c == 'C' {
			t[0], t[2] = t[2], t[0]
		}
	}
	for i := 0; i < 3; i++ {
		if t[i] == 1 {
			fmt.Println(i + 1)
			break
		}
	}
}
