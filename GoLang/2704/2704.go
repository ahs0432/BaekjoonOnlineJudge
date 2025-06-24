package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	var h, m, s int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		h, _ = strconv.Atoi(tmp[0:2])
		m, _ = strconv.Atoi(tmp[3:5])
		s, _ = strconv.Atoi(tmp[6:])

		for j := 5; j >= 0; j-- {
			if h&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}

			if m&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}

			if s&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}
		}

		fmt.Fprint(writer, " ")
		for j := 5; j >= 0; j-- {
			if h&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}
		}

		for j := 5; j >= 0; j-- {
			if m&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}
		}

		for j := 5; j >= 0; j-- {
			if s&(1<<j) >= 1 {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
