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

	sum := int64(0)
	for i := int64(1); i < n; i++ {
		sum += n*i + i
	}
	fmt.Println(sum)
}
