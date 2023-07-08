package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, p int
	fmt.Fscanln(reader, &a, &p)

	a *= 7
	p *= 13

	if a < p {
		fmt.Println("Petra")
	} else if a > p {
		fmt.Println("Axel")
	} else {
		fmt.Println("lika")
	}
}
