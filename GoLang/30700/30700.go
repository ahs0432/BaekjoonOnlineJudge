package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	check := "KOREA"
	now := 0

	for _, c := range s {
		if check[now%5] == byte(c) {
			now++
		}
	}
	fmt.Println(now)
}
