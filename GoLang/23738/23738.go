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

	new := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			new = append(new, 'v')
		} else if s[i] == 'E' {
			new = append(new, 'y')
			new = append(new, 'e')
		} else if s[i] == 'H' {
			new = append(new, 'n')
		} else if s[i] == 'P' {
			new = append(new, 'r')
		} else if s[i] == 'C' {
			new = append(new, 's')
		} else if s[i] == 'Y' {
			new = append(new, 'u')
		} else if s[i] == 'X' {
			new = append(new, 'h')
		} else {
			new = append(new, s[i]+32)
		}
	}
	fmt.Println(string(new))
}
