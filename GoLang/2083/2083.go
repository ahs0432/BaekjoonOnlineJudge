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
		var name string
		var old, kg int
		fmt.Fscanln(reader, &name, &old, &kg)

		if name == "#" && old == 0 && kg == 0 {
			break
		}

		if old > 17 || kg >= 80 {
			fmt.Fprintln(writer, name+" Senior")
		} else {
			fmt.Fprintln(writer, name+" Junior")
		}
	}
	writer.Flush()
}
