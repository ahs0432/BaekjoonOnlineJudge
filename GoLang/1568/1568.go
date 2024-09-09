package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	count := 0
	for k := int64(1); n > 0; {
		if k > n {
			k = 1
		}
		n -= k
		k++
		count++
	}

	fmt.Println(count)
}
