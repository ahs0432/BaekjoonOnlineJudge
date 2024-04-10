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

	t := []int{31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var y, m int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &y, &m)

		if m == 1 {
			fmt.Fprintln(writer, y-1, 12, 31)
		} else if m == 3 {
			if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
				fmt.Fprintln(writer, y, 2, 29)
			} else {
				fmt.Fprintln(writer, y, 2, 28)
			}
		} else {
			fmt.Fprintln(writer, y, m-1, t[m-2])
		}
	}
	writer.Flush()
}
