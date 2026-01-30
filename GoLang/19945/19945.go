package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, ans int
	fmt.Fscanln(reader, &n)

	if n < 0 {
		ans = 32
	} else {
		if n == 0 {
			ans = 1
		}

		for n != 0 {
			n >>= 1
			ans++
		}
	}
	fmt.Println(ans)
}
