package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &k)

	sum := n
	for i := 1; i <= k; i++ {
		tmp := 1
		for j := 0; j < i; j++ {
			tmp *= 10
		}

		sum += (n * tmp)
	}
	fmt.Println(sum)
}
