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

	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		b := []byte(s)

		if 97 <= b[0] && 122 >= b[0] {
			b[0] -= 32
		}

		fmt.Fprintln(writer, string(b[0:len(b)-1]))
	}

	writer.Flush()
}
