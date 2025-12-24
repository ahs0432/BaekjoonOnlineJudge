package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	fmt.Fprint(writer, "<")

	for len(queue) > 1 {
		for i := 1; i < k; i++ {
			tmp := queue[0]
			queue = queue[1:]
			queue = append(queue, tmp)
		}
		fmt.Fprintf(writer, "%d, ", queue[0])
		queue = queue[1:]
	}
	fmt.Fprintf(writer, "%d>\n", queue[0])
	writer.Flush()
}
