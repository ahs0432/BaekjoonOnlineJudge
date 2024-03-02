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

	tmp1 := ""
	for i := 0; i < n; i++ {
		tmp1 += "@@@"
	}
	for i := 0; i < n; i++ {
		tmp1 += " "
	}
	for i := 0; i < n; i++ {
		tmp1 += "@"
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, tmp1)
	}

	tmp2 := ""
	for i := 0; i < n; i++ {
		tmp2 += "@"
	}
	for i := 0; i < n; i++ {
		tmp2 += " "
	}
	for i := 0; i < n; i++ {
		tmp2 += "@"
	}
	for i := 0; i < n; i++ {
		tmp2 += " "
	}
	for i := 0; i < n; i++ {
		tmp2 += "@"
	}

	for i := 0; i < n*3; i++ {
		fmt.Fprintln(writer, tmp2)
	}

	tmp3 := ""
	for i := 0; i < n; i++ {
		tmp3 += "@"
	}
	for i := 0; i < n; i++ {
		tmp3 += " "
	}
	for i := 0; i < n; i++ {
		tmp3 += "@@@"
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, tmp3)
	}

	writer.Flush()
}
