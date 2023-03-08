package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, m, s int
	fmt.Fscanln(reader, &h, &m, &s)

	var add int
	fmt.Fscanln(reader, &add)

	h += (add / 3600)
	add %= 3600

	s += (add % 60)
	if s >= 60 {
		s -= 60
		m += 1
	}

	m += (add / 60)
	if m >= 60 {
		m -= 60
		h += 1
	}

	if h >= 24 {
		h %= 24
	}

	fmt.Println(h, m, s)
}
