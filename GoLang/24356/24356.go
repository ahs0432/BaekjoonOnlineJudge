package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t1, t2, m1, m2 int
	fmt.Fscanln(reader, &t1, &m1, &t2, &m2)

	m1 += t1 * 60
	m2 += t2 * 60

	if m1 > m2 {
		m2 += (24 * 60)
	}

	fmt.Println(m2-m1, (m2-m1)/30)
}
