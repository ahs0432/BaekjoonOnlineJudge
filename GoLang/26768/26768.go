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
	stmp := []byte(s)

	for i := 0; i < len(stmp); i++ {
		if stmp[i] == byte('a') {
			stmp[i] = byte('4')
		} else if s[i] == byte('e') {
			stmp[i] = byte('3')
		} else if s[i] == byte('i') {
			stmp[i] = byte('1')
		} else if s[i] == byte('o') {
			stmp[i] = byte('0')
		} else if s[i] == byte('s') {
			stmp[i] = byte('5')
		}
	}

	fmt.Println(string(stmp))
}
