package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func rev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func tail(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return s[1:]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var s string
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &s)

		dashSLen := int(math.Ceil(float64(len(s)) / 3.0))
		dashS := s[:dashSLen]

		revDashS := rev(dashS)
		tailDashS := tail(dashS)
		tailRevDashS := tail(revDashS)

		if s == (dashS+revDashS+dashS) ||
			s == (dashS+tailRevDashS+dashS) ||
			s == (dashS+revDashS+tailDashS) ||
			s == (dashS+tailRevDashS+tailDashS) {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
	writer.Flush()
}
