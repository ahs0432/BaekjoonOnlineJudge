package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	fmt.Fscanln(reader, &s)

	if s[0] == 'E' {
		fmt.Fprint(writer, "I")
	} else {
		fmt.Fprint(writer, "E")
	}

	if s[1] == 'N' {
		fmt.Fprint(writer, "S")
	} else {
		fmt.Fprint(writer, "N")
	}

	if s[2] == 'F' {
		fmt.Fprint(writer, "T")
	} else {
		fmt.Fprint(writer, "F")
	}

	if s[3] == 'J' {
		fmt.Fprint(writer, "P")
	} else {
		fmt.Fprint(writer, "J")
	}

	fmt.Fprintln(writer)
	writer.Flush()
}
