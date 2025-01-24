package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	var tmp []byte
	for {
		fmt.Fscanln(reader, &s)

		if s == "#" {
			break
		}

		tmp = []byte(s)

		for i := range tmp {
			if tmp[i] == 'a' || tmp[i] == 'e' || tmp[i] == 'i' || tmp[i] == 'o' || tmp[i] == 'u' {
				tmp = append(tmp[i:], tmp[:i]...)
				break
			}
		}

		tmp = append(tmp, 'a')
		tmp = append(tmp, 'y')

		fmt.Fprintln(writer, string(tmp))
	}
	writer.Flush()
}
