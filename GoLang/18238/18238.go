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

	sum, tmpR, tmpL, now := 0, 0, 0, 'A'
	for _, c := range s {
		tmpR = int(c - now)
		tmpL = int(now - c)
		if tmpR < 0 {
			tmpR += 26
		} else {
			tmpL += 26
		}

		if tmpR < tmpL {
			sum += tmpR
		} else {
			sum += tmpL
		}
		now = c
	}

	fmt.Println(sum)
}
