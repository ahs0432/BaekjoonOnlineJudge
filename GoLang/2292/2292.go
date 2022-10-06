package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	i := 1
	num -= 1

	for ; num > 0; i++ {
		num -= (6 * i)
	}

	fmt.Println(i)
}
