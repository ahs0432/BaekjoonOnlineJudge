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

	suma, sumb := 0, 0
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a > b {
			suma++
		} else if a < b {
			sumb++
		}
	}

	fmt.Println(suma, sumb)
}
