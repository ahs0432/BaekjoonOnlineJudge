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

	for i := 0; ; i++ {
		if n < 10 {
			fmt.Println(i)
			break
		}

		tmp := n % 10
		n /= 10
		for n != 0 {
			tmp *= (n % 10)
			n /= 10
		}
		n = tmp
	}
}
