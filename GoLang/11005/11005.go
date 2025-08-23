package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	var tmp int
	ans := ""
	for a != 0 {
		tmp = a % b
		if tmp >= 0 && tmp <= 9 {
			ans = string(byte(tmp+'0')) + ans
		} else {
			ans = string(byte(tmp-10+'A')) + ans
		}
		a /= b
	}
	fmt.Println(ans)
}
