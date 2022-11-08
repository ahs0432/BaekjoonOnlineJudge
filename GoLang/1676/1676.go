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

	answer := 0
	for n >= 5 {
		answer += (n / 5)
		n /= 5
	}

	fmt.Println(answer)
}
