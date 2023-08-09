package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, s, g, a, b, l, r float64
	fmt.Fscanln(reader, &m, &s, &g)
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &l, &r)

	lw := l / a
	rw := r / b

	ls := m / g
	if int(m)%int(g) == 1 {
		ls += 1
	}

	rs := m / s
	if int(m)%int(s) == 1 {
		ls += 1
	}

	if ls < rs {
		if ls+lw < rs+rw {
			fmt.Println("friskus")
		} else {
			fmt.Println("latmask")
		}
	} else {
		if ls+lw < rs+rw {
			fmt.Println("friskus")
		} else {
			fmt.Println("latmask")
		}
	}

}
