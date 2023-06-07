package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var p float64
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &p)

	tmp := p

	if n >= 5 && tmp > p-500.0 {
		tmp = p - 500.0
	}

	if n >= 10 && tmp > p*0.9 {
		tmp = p * 0.9
	}

	if n >= 15 && tmp > p-2000.0 {
		tmp = p - 2000.0
	}

	if n >= 20 && tmp > p*0.75 {
		tmp = p * 0.75
	}

	if tmp < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(tmp)
	}
}
