package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &s)

	alpha := make([]int, 26)

	for i := 0; i < len(s); i++ {
		if int(s[i]) > 96 {
			alpha[s[i]-97]++
		} else {
			alpha[s[i]-65]++
		}
	}

	max := alpha[0]
	top := "A"

	for i := 1; i < len(alpha); i++ {
		if max < alpha[i] {
			top = string(rune(i + 65))
			max = alpha[i]
		} else if max == alpha[i] {
			top = "?"
		}
	}

	fmt.Println(top)
}
