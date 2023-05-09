package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, m, l int
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &l)

	if s+m*2+l*3 >= 10 {
		fmt.Println("happy")
	} else {
		fmt.Println("sad")
	}
}
