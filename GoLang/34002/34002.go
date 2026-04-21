package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c, s, v, l int
	fmt.Fscanln(reader, &a, &b, &c)
	fmt.Fscanln(reader, &s, &v)
	fmt.Fscanln(reader, &l)

	ceil := func(x, y int) int {
		return (x + y - 1) / y
	}

	t := 25000 - 100*l

	if t <= 30*c*v {
		fmt.Fprintln(writer, ceil(t, c))
	} else if t-30*c*v <= 30*b*s {
		fmt.Fprintln(writer, ceil(t-30*c*v, b)+30*v)
	} else {
		fmt.Fprintln(writer, ceil(t-30*c*v-30*b*s, a)+30*(s+v))
	}
}
