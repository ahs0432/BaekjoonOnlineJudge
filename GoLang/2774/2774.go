package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var x string
	var count int
	var check []bool
	for i := 0; i < t; i++ {
		count = 0
		check = make([]bool, 10)
		fmt.Fscanln(reader, &x)

		for i := 0; i < len(x); i++ {
			check[x[i]-48] = true
		}

		for i := 0; i < 10; i++ {
			if check[i] {
				count++
			}
		}
		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
