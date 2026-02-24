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

	now := []byte{s[0]}
	count := 0

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i]-1 {
			now = append(now, s[i])
		} else {
			if len(now) == 3 {
				count++
			}
			now = []byte{s[i]}
		}
	}
	if len(now) == 3 {
		count++
	}
	fmt.Println(count)
}
