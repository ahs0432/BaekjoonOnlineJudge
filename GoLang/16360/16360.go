package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var str []byte
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str)

		if str[len(str)-1] == 'a' {
			str = append(str, 's')
		} else if str[len(str)-1] == 'i' || str[len(str)-1] == 'y' {
			str[len(str)-1] = 'i'
			str = append(str, 'o')
			str = append(str, 's')
		} else if str[len(str)-1] == 'l' {
			str = append(str, 'e')
			str = append(str, 's')
		} else if str[len(str)-1] == 'n' {
			str[len(str)-1] = 'a'
			str = append(str, 'n')
			str = append(str, 'e')
			str = append(str, 's')
		} else if str[len(str)-1] == 'e' && str[len(str)-2] == 'n' {
			str[len(str)-2] = 'a'
			str[len(str)-1] = 'n'
			str = append(str, 'e')
			str = append(str, 's')
		} else if str[len(str)-1] == 'o' {
			str = append(str, 's')
		} else if str[len(str)-1] == 'r' {
			str = append(str, 'e')
			str = append(str, 's')
		} else if str[len(str)-1] == 't' {
			str = append(str, 'a')
			str = append(str, 's')
		} else if str[len(str)-1] == 'u' {
			str = append(str, 's')
		} else if str[len(str)-1] == 'v' {
			str = append(str, 'e')
			str = append(str, 's')
		} else if str[len(str)-1] == 'w' {
			str = append(str, 'a')
			str = append(str, 's')
		} else {
			str = append(str, 'u')
			str = append(str, 's')
		}
		fmt.Fprintln(writer, string(str))
	}
	writer.Flush()
}
