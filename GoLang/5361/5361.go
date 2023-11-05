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

	var a, b, c, d, e float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &c, &d, &e)
		fmt.Fprintf(writer, "$%.2f\n", a*350.34+b*230.90+c*190.55+d*125.30+e*180.90)
	}
	writer.Flush()
}
