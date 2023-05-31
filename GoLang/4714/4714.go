package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var tmp float64
		fmt.Fscanln(reader, &tmp)

		if tmp == -1.0 {
			break
		}

		fmt.Fprintf(writer, "Objects weighing %.2f on Earth will weigh %.2f on the moon.\n", tmp, tmp*0.167)
	}
	writer.Flush()
}
