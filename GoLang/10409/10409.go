package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscanln(reader, &n, &t)

	i := 0
	var tmp int
	for ; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		t -= tmp

		if t < 0 {
			break
		}
	}

	fmt.Println(i)
}
