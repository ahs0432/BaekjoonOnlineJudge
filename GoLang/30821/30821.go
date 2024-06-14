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

	var count int64 = 1

	for i := int64(0); i < 5; i++ {
		count *= n - i
	}

	for i := int64(0); i < 5; i++ {
		count /= (i + 1)
	}

	fmt.Println(count)
}
