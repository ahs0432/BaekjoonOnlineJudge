package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, w, m int
	fmt.Fscanln(reader, &k, &w, &m)

	sum := (w - k) / m

	if (w-k)%m != 0 {
		sum++
	}

	fmt.Println(sum)
}
