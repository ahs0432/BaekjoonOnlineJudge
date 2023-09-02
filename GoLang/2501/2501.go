package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	cnt := 0
	p := true
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}

		if cnt == k {
			p = false
			fmt.Println(i)
			break
		}
	}

	if p {
		fmt.Println(0)
	}
}
