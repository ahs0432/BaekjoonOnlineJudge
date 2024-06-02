package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := 1
	var tmp int
	for {
		fmt.Println("? A", a)
		fmt.Fscanln(reader, &tmp)
		if tmp == 1 {
			break
		}
		a++
	}

	b := 1
	tmp = 0
	for {
		fmt.Println("? B", b)
		fmt.Fscanln(reader, &tmp)
		if tmp == 1 {
			break
		}
		b++
	}
	fmt.Println("!", a+b)
}
