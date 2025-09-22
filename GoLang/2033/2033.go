package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	count := 0
	for n/10 > 0 {
		if n%10 > 4 {
			n += 10
		}
		n /= 10
		count++
	}

	for i := 0; i < count; i++ {
		n *= 10
	}
	fmt.Println(n)
}
