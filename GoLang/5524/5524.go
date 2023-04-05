package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, strings.ToLower(tmp))
	}

	writer.Flush()
}
