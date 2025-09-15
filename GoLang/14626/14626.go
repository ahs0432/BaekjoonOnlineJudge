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

	sum := 0
	sCheck := false
	nCheck := false
	for i := 0; i < 13; i++ {
		if i%2 == 0 {
			nCheck = false
		} else {
			nCheck = true
		}

		if s[i] == '*' {
			sCheck = nCheck
		} else {
			if nCheck {
				sum += int(s[i]-'0') * 3
			} else {
				sum += int(s[i] - '0')
			}
		}
	}

	sum %= 10
	if sCheck {
		for i := 0; i < 10; i++ {
			if (sum+i*3)%10 == 0 {
				fmt.Println(i)
				break
			}
		}
	} else {
		fmt.Println((10 - sum) % 10)
	}
}
