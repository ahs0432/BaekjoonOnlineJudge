package main

import (
	"bufio"
	"fmt"
	"os"
)

func tCalc(n int) int {
	return (n*n + n) / 2
}

func check(count int, arrNum int) int {
	for a := 1; a <= count; a++ {
		for b := 1; b <= count; b++ {
			for c := 1; c <= count; c++ {
				if arrNum == (tCalc(a) + tCalc(b) + tCalc(c)) {
					return 1
				}
			}
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	arr := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &arr[i])
	}

	for _, val := range arr {
		count := 0
		tVal := 0
		for tVal < val {
			count++
			tVal = tCalc(count)
		}
		fmt.Fprintln(writer, check(count, val))
	}
	writer.Flush()
}
