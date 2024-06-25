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

	var k, n float64
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &k, &n)

		s1 := int(n * (n + 1) / 2)
		s2 := int((n * 2) * (n / 2))
		s3 := int((n*2 + 2) * (n / 2))
		fmt.Fprintln(writer, k, s1, s2, s3)
	}
	writer.Flush()
}
