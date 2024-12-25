package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, n string
	fmt.Fscanln(reader, &h)
	fmt.Fscanln(reader, &n)

	ans := 0
	flag := true
	for i := 0; i < len(h); i++ {
		flag = true
		for j := 0; j < len(n); j++ {
			if len(h) <= i+j || h[i+j] != n[j] {
				flag = false
				break
			}
		}

		if flag {
			ans++
		}
	}
	fmt.Println(ans)
}
