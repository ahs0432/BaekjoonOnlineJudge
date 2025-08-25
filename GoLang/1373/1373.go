package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var oct string
	fmt.Fscanln(reader, &oct)

	for len(oct)%3 != 0 {
		oct = "0" + oct
	}

	for i := 0; i < len(oct); i += 3 {
		fmt.Fprint(writer, (oct[i]-'0')*4+(oct[i+1]-'0')*2+(oct[i+2]-'0'))
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
