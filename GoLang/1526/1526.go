package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var b bool
	for {
		b = true
		for _, c := range strconv.Itoa(n) {
			if c != '4' && c != '7' {
				b = false
				n -= 1
			}
		}

		if b {
			fmt.Println(n)
			break
		}
	}
}
