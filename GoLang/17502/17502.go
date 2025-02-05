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

	var s string
	fmt.Fscanln(reader, &s)

	t := []byte(s)
	for i := 0; i < len(s)/2+(n%2); i++ {
		if t[i] == '?' {
			if t[len(s)-1-i] == '?' {
				t[len(s)-1-i] = 'a'
			}
			t[i] = t[len(s)-1-i]
		} else if t[len(s)-1-i] == '?' {
			if t[i] == '?' {
				t[i] = 'a'
			}
			t[len(s)-1-i] = t[i]
		}
	}
	fmt.Println(string(t))
}
