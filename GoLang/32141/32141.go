package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, h int
	fmt.Fscanln(reader, &n, &h)

	var d int
	now, ans := 0, 0

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d)

		if now < h {
			ans++
			now += d

			if now >= h {
				break
			}
		}
	}

	if now < h {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
