package main

import (
	"bufio"
	"fmt"
	"os"
)

func f2017(n int) int {
	if n == 1 {
		return 5000000
	} else if n >= 2 && n <= 3 {
		return 3000000
	} else if n >= 4 && n <= 6 {
		return 2000000
	} else if n >= 7 && n <= 10 {
		return 500000
	} else if n >= 11 && n <= 15 {
		return 300000
	} else if n >= 16 && n <= 21 {
		return 100000
	}
	return 0
}

func f2018(n int) int {
	if n == 1 {
		return 5120000
	} else if n >= 2 && n <= 3 {
		return 2560000
	} else if n >= 4 && n <= 7 {
		return 1280000
	} else if n >= 8 && n <= 15 {
		return 640000
	} else if n >= 16 && n <= 31 {
		return 320000
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, f2017(a)+f2018(b))
	}
	writer.Flush()
}
