package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var y1, m1, d1 int
	var y2, m2, d2 int
	fmt.Fscanln(reader, &y1, &m1, &d1)
	fmt.Fscanln(reader, &y2, &m2, &d2)

	yearOld := y2 - y1
	old := yearOld

	if m2 < m1 {
		old -= 1
	} else if m2 == m1 && d2 < d1 {
		old -= 1
	}

	fmt.Println(old)
	fmt.Println(yearOld + 1)
	fmt.Println(yearOld)
}
