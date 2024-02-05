package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var k int
	fmt.Fscanln(reader, &k)

	var h int
	var c string
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &h)
		fmt.Fscanln(reader, &c)

		for _, tmp := range c {
			if tmp == 'c' {
				h++
			} else {
				h--
			}
		}
		fmt.Fprintf(writer, "Data Set %d:\n%d\n\n", i+1, h)
	}
	writer.Flush()
}
