package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var a, b int
	var s string
	var ans []byte
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fscanln(reader, &s)

		ans = []byte{}
		for _, c := range s {
			ans = append(ans, byte((((a*int(c-'A'))+b)%26))+'A')
		}
		fmt.Fprintln(writer, string(ans))
	}
	writer.Flush()
}
