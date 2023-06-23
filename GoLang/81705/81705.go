package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p, q int
	fmt.Fscanln(reader, &p)
	fmt.Fscanln(reader, &q)

	if p <= 50 && q <= 10 {
		fmt.Println("White")
	} else if q > 30 {
		fmt.Println("Red")
	} else {
		fmt.Println("Yellow")
	}
}
