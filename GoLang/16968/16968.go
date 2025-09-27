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

	ans := 1
	if s[0] == 'd' {
		ans = 10
	} else {
		ans = 26
	}

	for i := 1; i < len(s); i++ {
		if s[i] == 'd' {
			if s[i-1] == 'd' {
				ans *= 9
			} else {
				ans *= 10
			}
		} else {
			if s[i] == 'c' {
				if s[i-1] == 'c' {
					ans *= 25
				} else {
					ans *= 26
				}
			}
		}
	}
	fmt.Println(ans)
}
