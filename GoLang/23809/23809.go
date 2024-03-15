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

	s1 := ""
	for i := 0; i < n; i++ {
		s1 += "@"
	}
	for i := 0; i < n; i++ {
		s1 += "   "
	}
	for i := 0; i < n; i++ {
		s1 += "@"
	}

	s2 := ""
	for i := 0; i < n; i++ {
		s2 += "@"
	}
	for i := 0; i < n; i++ {
		s2 += "  "
	}
	for i := 0; i < n; i++ {
		s2 += "@"
	}

	s3 := ""
	for i := 0; i < n*3; i++ {
		s3 += "@"
	}

	for i := 0; i < n; i++ {
		fmt.Println(s1)
	}
	for i := 0; i < n; i++ {
		fmt.Println(s2)
	}
	for i := 0; i < n; i++ {
		fmt.Println(s3)
	}
	for i := 0; i < n; i++ {
		fmt.Println(s2)
	}
	for i := 0; i < n; i++ {
		fmt.Println(s1)
	}

	writer.Flush()
}
