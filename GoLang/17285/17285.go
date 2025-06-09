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

	tmp := s[0] ^ 'C'
	str := []byte{}

	for i := 0; i < len(s); i++ {
		str = append(str, s[i]^tmp)
	}
	fmt.Println(string(str))
}
