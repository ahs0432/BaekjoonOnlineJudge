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

	ans := 0
	for h := 0; h <= n; h++ {
		for m := 0; m <= 59; m++ {
			for s := 0; s <= 59; s++ {
				if h%10 == k || h/10 == k || m%10 == k || m/10 == k || s%10 == k || s/10 == k {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
