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

	var m, tmpi int
	var s, tmps string
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &m, &s)
		for j := 0; j < m; j++ {
			if s == "C" {
				if j == m-1 {
					fmt.Fscanln(reader, &tmps)
				} else {
					fmt.Fscan(reader, &tmps)
				}
				fmt.Fprint(writer, int(tmps[0]+1-'A'), " ")
			} else {
				if j == m-1 {
					fmt.Fscanln(reader, &tmpi)
				} else {
					fmt.Fscan(reader, &tmpi)
				}
				fmt.Fprint(writer, string(rune(tmpi-1+'A')), " ")
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
