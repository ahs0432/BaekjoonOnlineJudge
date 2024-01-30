package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var d, r, t float64

	i := 1
	for {
		fmt.Fscanln(reader, &d, &r, &t)
		if r == 0 {
			break
		}

		fmt.Fprintf(writer, "Trip #%d: %.2f %.2f\n", i, d/63360*3.1415927*r, (d/63360*3.1415927*r)/t*3600)
		i += 1
	}
	writer.Flush()
}
