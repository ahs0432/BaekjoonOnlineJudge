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

	if s[0] == '0' {
		sum := 0
		n := 1
		if len(s) >= 2 && s[1] == 'x' {
			for i := len(s) - 1; i > 1; i-- {
				if s[i] >= '0' && s[i] <= '9' {
					sum += (int(s[i]-48) * n)
				} else {
					sum += (int(s[i]-87) * n)
				}
				n *= 16
			}
			fmt.Println(sum)
		} else {
			for i := len(s) - 1; i > 0; i-- {
				sum += (int(s[i]-48) * n)
				n *= 8
			}
			fmt.Println(sum)
		}
	} else {
		fmt.Println(s)
	}
}
