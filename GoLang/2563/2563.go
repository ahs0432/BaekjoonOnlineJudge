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

	var table [100][100]int

	var sx, sy int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &sx, &sy)
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				table[sy+y][sx+x] = 1
			}
		}
	}

	count := 0
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			count += table[y][x]
		}
	}
	fmt.Println(count)
}
