package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t float64

	for {
		fmt.Fscanln(reader, &t)
		if t == 0.0 {
			break
		}

		tmp := 0.0
		i := 2

		for tmp < t {
			tmp += 1.0 / float64(i)
			i++
		}
		fmt.Fprintf(writer, "%d card(s)\n", i-2)
	}
	writer.Flush()
}
