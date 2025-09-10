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

	var a, b, x, y int
	arr := make([][]int, 1001)
	for i := 0; i < 1001; i++ {
		arr[i] = make([]int, 1001)
	}

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &x, &y)
		for j := a; j < a+x; j++ {
			for k := b; k < b+y; k++ {
				arr[j][k] = i + 1
			}
		}
	}

	var count int
	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < 1001; j++ {
			for k := 0; k < 1001; k++ {
				if arr[j][k] == i+1 {
					count++
				}
			}
		}
		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
