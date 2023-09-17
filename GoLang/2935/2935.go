package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, s string

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &b)

	if s == "*" {
		fmt.Println(a + b[1:])
	} else {
		var sum []byte
		if len(a) > len(b) {
			sum = []byte(a)
			sum[len(a)-len(b)] = byte('1')
		} else if len(a) == len(b) {
			sum = []byte(a)
			sum[0] = byte('2')
		} else {
			sum = []byte(b)
			sum[len(b)-len(a)] = byte('1')
		}
		fmt.Println(string(sum))
	}
}
