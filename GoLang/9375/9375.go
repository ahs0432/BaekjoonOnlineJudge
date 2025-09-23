package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var c int
		fmt.Fscanln(reader, &c)

		wears := map[string]int{}

		for j := 0; j < c; j++ {
			var name, types string
			fmt.Fscanln(reader, &name, &types)

			if _, e := wears[types]; !e {
				wears[types] = 1
			} else {
				wears[types]++
			}
		}

		answer := 1
		for _, v := range wears {
			answer *= (v + 1)
		}
		fmt.Fprintln(writer, answer-1)
	}
	writer.Flush()
}
