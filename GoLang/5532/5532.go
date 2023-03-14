package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, a, b, c, d int
	fmt.Fscanln(reader, &l)
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &d)

	s1 := (a / c)
	s2 := (b / d)

	if (a % c) >= 1 {
		s1 += 1
	}

	if (b % d) >= 1 {
		s2 += 1
	}

	if s1 < s2 {
		s1 = s2
	}

	fmt.Println(l - s1)
}
