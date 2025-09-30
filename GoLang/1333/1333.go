package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l, d int
	fmt.Fscanln(reader, &n, &l, &d)

	time := 0
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			time++
		}

		for j := 0; j < 5; j++ {
			if time%d == 0 {
				ans = time
				i = n
				break
			}
			time++
		}

		if i == n-1 && ans == 0 {
			for {
				if time%d == 0 {
					ans = time
					break
				}
				time++
			}
		}
	}

	fmt.Println(ans)
}
