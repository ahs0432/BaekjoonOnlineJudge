package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, m int
	fmt.Fscanln(reader, &a, &m)

	for i := 1; i <= 10000; i++ {
		if a*i%m == 1 {
			fmt.Println(i)
			break
		}
	}
}
