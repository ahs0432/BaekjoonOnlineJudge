package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -(a)
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	t := make([]int, 4)
	c := 0
	for {
		fmt.Fscanln(reader, &t[0], &t[1], &t[2], &t[3])

		if t[0] == 0 && t[1] == 0 && t[2] == 0 && t[3] == 0 {
			break
		}

		c = 0
		for !(t[0] == t[1] && t[1] == t[2] && t[2] == t[3]) {
			t[0], t[1], t[2], t[3] = abs(t[0]-t[1]), abs(t[1]-t[2]), abs(t[2]-t[3]), abs(t[3]-t[0])
			c++
		}
		fmt.Fprintln(writer, c)
	}
	writer.Flush()
}
