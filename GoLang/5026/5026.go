package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp == "P=NP" {
			fmt.Fprintln(writer, "skipped")
			continue
		}

		a, _ = strconv.Atoi(strings.Split(tmp, "+")[0])
		b, _ = strconv.Atoi(strings.Split(tmp, "+")[1])
		fmt.Fprintln(writer, a+b)
	}
	writer.Flush()
}
