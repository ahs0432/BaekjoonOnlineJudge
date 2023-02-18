package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	for i := 0; i < 5; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		sum += n
	}
	fmt.Println(sum)
}
