package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var tmp int
	for i := 0; ; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp == 0 {
			break
		}

		if tmp%2 == 0 {
			tmp = (((tmp * 3) / 2) * 3) / 9
			fmt.Fprintf(writer, "%d. even %d\n", i+1, tmp)
		} else {
			tmp = ((((tmp * 3) + 1) / 2) * 3) / 9
			fmt.Fprintf(writer, "%d. odd %d\n", i+1, tmp)
		}
	}

	writer.Flush()
}
