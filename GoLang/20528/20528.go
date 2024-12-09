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

	var tmp string
	fmt.Fscan(reader, &tmp)
	now := tmp[0]

	ans := 1
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &tmp)
		if now != tmp[0] {
			ans = 0
			break
		}
	}

	fmt.Println(ans)
}
