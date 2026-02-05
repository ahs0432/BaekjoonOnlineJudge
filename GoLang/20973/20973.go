package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, tmp string
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &tmp)

	sLen, tmpLen := len(s), len(tmp)
	c1, c2 := 0, 0
	ans := 1
	for {
		if c2 == tmpLen {
			break
		}

		if c1 == sLen {
			c1 = 0
			ans++
		}

		if s[c1] == tmp[c2] {
			c1++
			c2++
		} else {
			c1++
		}
	}
	fmt.Println(ans)
}
