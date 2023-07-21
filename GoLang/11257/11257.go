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

	for i := 0; i < n; i++ {
		var a int
		var s1, s2, s3 float64
		fmt.Fscanln(reader, &a, &s1, &s2, &s3)

		sum := s1 + s2 + s3
		fmt.Fprint(writer, a, sum, " ")

		if sum < 55 {
			fmt.Fprintln(writer, "FAIL")
		} else if s1 < (35.0*0.3) || s2 < (25.0*0.3) || s3 < (40.0*0.3) {
			fmt.Fprintln(writer, "FAIL")
		} else {
			fmt.Fprintln(writer, "PASS")
		}
	}
	writer.Flush()
}
