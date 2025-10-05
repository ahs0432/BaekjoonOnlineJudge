package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(n int64) int64 {
	s := ""
	for n > 0 {
		if n%2 == 1 {
			s += "1"
		} else {
			s += "0"
		}
		n /= 2
	}

	ans := int64(0)
	for i := 0; i < len(s); i++ {
		ans = ans*2 + int64(s[i]-'0')
	}
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)
	fmt.Println(check(n))

}
