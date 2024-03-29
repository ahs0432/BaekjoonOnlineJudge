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

	var tmp int
	max := 0
	chk := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if i != 0 {
			if max < tmp {
				max = tmp
				chk = 1
			}
		} else {
			max = tmp
		}
	}

	fmt.Println(string("SN"[chk]))
}
