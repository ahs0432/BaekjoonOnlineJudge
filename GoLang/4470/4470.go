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

	for i := 1; i <= n; i++ {
		tmp, _ := reader.ReadString('\n')
		fmt.Fprintf(writer, "%d. "+tmp, i)
	}
	writer.Flush()
}
