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
	n = (^n + 1) ^ n

	ans := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			ans++
		}
	}
	fmt.Println(ans)

}
