package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var k int
	fmt.Fscanln(reader, &k)

	var data []int
	var p, m int
	var tmp, count int
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &p, &m)
		data = make([]int, m+1)

		count = 0
		for j := 0; j < p; j++ {
			fmt.Fscanln(reader, &tmp)

			if data[tmp] == 0 {
				data[tmp] = 1
			} else {
				count++
			}
		}
		fmt.Fprintln(writer, count)
	}

	writer.Flush()
}
