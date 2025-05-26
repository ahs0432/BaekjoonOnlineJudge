package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	ans := 0
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				ans++
			} else {
				ans += 2
			}
		}
	}
	fmt.Println(ans)
}
