package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var p1, p2, e1, e2 int
	fmt.Fscanln(reader, &p1, &e2)
	fmt.Fscanln(reader, &e1, &p2)

	p := p1 + p2
	e := e1 + e2

	if p > e {
		fmt.Println("Persepolis")
	} else if p < e {
		fmt.Println("Esteghlal")
	} else {
		if p1 < e1 {
			fmt.Println("Persepolis")
		} else if p1 > e1 {
			fmt.Println("Esteghlal")
		} else {
			fmt.Println("Penalty")
		}
	}
}
